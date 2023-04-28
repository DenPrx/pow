package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"pow/internal/pow"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	conn, err := waitServerConnection()
	if err != nil {
		logger.Fatal("error connecting to server: ", err)
	}
	defer conn.Close()

	var challenge pow.Challenge
	err = binary.Read(conn, binary.BigEndian, &challenge)
	if err != nil {
		logger.Fatal("error receiving challenge from server: ", err)
	}

	solution, ok := pow.SolveChallenge(challenge)
	if !ok {
		logger.Fatal("failed solution doesn't found")
	}

	err = binary.Write(conn, binary.BigEndian, &solution)
	if err != nil {
		logger.Fatal("error sending solution to server: ", err)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		logger.Info(scanner.Text())
	}
}

func waitServerConnection() (net.Conn, error) {
	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", os.Getenv("SERVER_URL"))
		if err == nil {
			return conn, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("server didn't respond after 10 seconds")
}
