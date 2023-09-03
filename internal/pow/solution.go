package pow

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func Solve(challenge string) string {
	prefix := strings.Repeat("0", difficulty)
	i := 0
	for {
		data := challenge + fmt.Sprintf("%d", i)
		hash := sha256.Sum256([]byte(data))
		hashStr := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashStr, prefix) {
			return fmt.Sprintf("%d", i)
		}
		i++
	}
}
