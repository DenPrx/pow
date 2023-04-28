package pow

import (
	"testing"
)

func TestVerifySolution(t *testing.T) {
	challenge := Challenge{
		Seed:       123456789,
		Difficulty: 5,
	}
	solution := Solution{
		Nonce: 4,
	}

	if !VerifySolution(challenge, solution) {
		t.Error("solution is not valid")
	}
}

func TestSolveChallenge(t *testing.T) {
	challenge := Challenge{
		Seed:       123456789,
		Difficulty: 5,
	}

	solution, ok := SolveChallenge(challenge)
	if !ok {
		t.Error("challenge not passed")
	}

	if solution.Nonce == 0 {
		t.Error("nonce should not be zero")
	}

	if !VerifySolution(challenge, solution) {
		t.Error("solution is not valid")
	}
}
