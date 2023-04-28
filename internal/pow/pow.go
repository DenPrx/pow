package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
	"math/rand"
)

const hashSize = 256

type Challenge struct {
	Seed       uint64
	Difficulty int64
}

type Solution struct {
	Nonce int64
}

func GenerateChallenge(difficulty int64) Challenge {
	return Challenge{
		Seed:       generateSeed(),
		Difficulty: difficulty,
	}
}

func SolveChallenge(challenge Challenge) (Solution, bool) {
	var solution Solution
	for nonce := int64(0); nonce < math.MaxInt64; nonce++ {
		solution.Nonce = nonce

		headers := bytes.Join([][]byte{
			toBytes(challenge.Seed),
			toBytes(challenge.Difficulty),
			toBytes(solution.Nonce),
		}, []byte{})

		hash := sha256.Sum256(headers)

		target := big.NewInt(1)
		target.Lsh(target, uint(hashSize-challenge.Difficulty))
		hashInt := new(big.Int).SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			return solution, true
		}
	}

	return Solution{}, false
}

func VerifySolution(challenge Challenge, solution Solution) bool {
	headers := bytes.Join([][]byte{
		toBytes(challenge.Seed),
		toBytes(challenge.Difficulty),
		toBytes(solution.Nonce),
	}, []byte{})

	hash := sha256.Sum256(headers)

	target := big.NewInt(1)
	target.Lsh(target, uint(256-challenge.Difficulty))
	hashInt := new(big.Int).SetBytes(hash[:])
	return hashInt.Cmp(target) == -1
}

func generateSeed() uint64 {
	seed := rand.Uint64()
	return seed
}

func toBytes[V int64 | uint64](n V) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}
