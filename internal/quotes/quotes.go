package quotes

import (
	"bufio"
	"math/rand"
	"os"
)

type Quotes []string

func New(filename string) (Quotes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			quotes = append(quotes, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q Quotes) GetQuote() string {
	quote := q[rand.Intn(len(q))]

	return quote
}
