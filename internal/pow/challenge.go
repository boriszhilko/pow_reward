package pow

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	difficulty = 4
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var WordsOfWisdom = []string{
	"The only true wisdom is in knowing you know nothing.",
	"Turn your wounds into wisdom.",
	"Wisdom is the daughter of experience.",
}

func GetChallenge(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func VerifySolution(challenge, response string) bool {
	hash := sha256.Sum256([]byte(challenge + response))
	hashStr := fmt.Sprintf("%x", hash)
	isCorrect := strings.HasPrefix(hashStr, strings.Repeat("0", difficulty))
	return isCorrect
}

func GetReward() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	reward := WordsOfWisdom[r.Intn(len(WordsOfWisdom))]
	return reward
}
