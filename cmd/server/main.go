package main

import (
	"net"
	"os"
	"strconv"
	"time"

	"pow/internal/server"

	"github.com/sirupsen/logrus"
)

const defaultDifficulty = 10

func main() {
	logger := logrus.New()

	quotesPath := os.Getenv("QUOTES_PATH")
	difficultyStr := os.Getenv("DIFFICULTY")
	difficulty, err := strconv.ParseInt(difficultyStr, 10, 64)
	if err != nil {
		difficulty = defaultDifficulty
		logger.Warnf("difficulty didn't parsed: %s", difficultyStr)
	}

	serv, err := server.New(logger, quotesPath, difficulty)
	if err != nil {
		logger.Fatal("failed server initialisation: ", err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		logger.Fatal("failed starting listener: ", err)
	}

	logger.Info("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("error accepting connection: ", err)
			continue
		}
		err = conn.SetDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			logger.Error("error setting deadline: ", err)
			continue
		}

		go serv.HandleConnection(conn)
	}
}
