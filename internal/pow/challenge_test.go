package pow

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetChallenge(t *testing.T) {
	tests := []struct {
		length    int
		expectLen int
	}{
		{5, 5},
		{10, 10},
		{0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("GetChallenge of length %d", test.length), func(t *testing.T) {
			result := GetChallenge(test.length)
			if len(result) != test.expectLen {
				t.Errorf("Expected length %d, but got %d", test.expectLen, len(result))
			}
		})
	}
}

func TestVerifySolution(t *testing.T) {
	tests := []struct {
		challenge string
		response  string
		expected  bool
	}{
		{"abcd", "1234", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("VerifySolution with challenge %s and response %s", test.challenge, test.response), func(t *testing.T) {
			result := VerifySolution(test.challenge, test.response)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestGetReward(t *testing.T) {
	reward := GetReward()
	if !strings.Contains(strings.Join(WordsOfWisdom, " "), reward) {
		t.Errorf("Expected reward to be one of the WISDOMs, but got %s", reward)
	}
}
