package server

import (
	"encoding/binary"
	"fmt"
	"net"

	"pow/internal/pow"
	"pow/internal/quotes"

	"github.com/sirupsen/logrus"
)

type server struct {
	Logger     *logrus.Logger
	Quotes     quotes.Quotes
	Difficulty int64
}

func New(logger *logrus.Logger, quotesFilename string, difficulty int64) (server, error) {
	wisdomQuotes, err := quotes.New(quotesFilename)
	if err != nil {
		return server{}, err
	}

	return server{
		Logger:     logger,
		Quotes:     wisdomQuotes,
		Difficulty: difficulty,
	}, nil
}

func (s server) HandleConnection(conn net.Conn) {
	defer conn.Close()

	s.Logger.Info("client connected")

	challenge := pow.GenerateChallenge(s.Difficulty)

	err := binary.Write(conn, binary.BigEndian, &challenge)
	if err != nil {
		s.Logger.Error("error sending challenge to client:", err)
		return
	}

	var solution pow.Solution
	err = binary.Read(conn, binary.BigEndian, &solution)
	if err != nil {
		s.Logger.Error("error receiving solution from client:", err)
		return
	}

	s.Logger.Infof("sollution: %v", solution.Nonce)

	if pow.VerifySolution(challenge, solution) {
		quote := s.Quotes.GetQuote()
		_, err = fmt.Fprintf(conn, "%s\n", quote)
		if err != nil {
			s.Logger.Error("error sending quote to client:", err)
			return
		}
		s.Logger.Info("PoW verification passed")
	} else {
		s.Logger.Warn("PoW verification failed")
	}
}
