package pow

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		challenge string
	}{
		{"abcd"},
		{"test1234"},
		{"hello"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Solve challenge %s", tt.challenge), func(t *testing.T) {
			solution := Solve(tt.challenge)

			data := tt.challenge + solution
			hash := sha256.Sum256([]byte(data))
			hashStr := fmt.Sprintf("%x", hash)
			prefix := strings.Repeat("0", difficulty)

			if !strings.HasPrefix(hashStr, prefix) {
				t.Errorf("Solve function returned an incorrect solution for challenge %s", tt.challenge)
			}
		})
	}
}
