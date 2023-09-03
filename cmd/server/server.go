package main

import (
	"fmt"
	"net"
	"os"
	"pow_reward/internal/pow"
)

const BufferSize = 256

func main() {
	listener := createServer(":8080")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}

func createServer(address string) net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error listening on %s: %v\n", address, err)
		os.Exit(1)
	}
	fmt.Printf("Listening on %s\n", address)
	return listener
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := sendChallenge(conn)
	response := receiveResponse(conn)

	isCorrect := pow.VerifySolution(challenge, response)
	sendFeedback(conn, isCorrect)
}

func sendChallenge(conn net.Conn) string {
	challenge := pow.GetChallenge(10)
	conn.Write([]byte(challenge))
	return challenge
}

func receiveResponse(conn net.Conn) string {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return ""
	}
	return string(buffer[:n])
}

func sendFeedback(conn net.Conn, isCorrect bool) {
	if isCorrect {
		sendRandomReward(conn)
	} else {
		conn.Write([]byte("Incorrect PoW"))
	}
}

func sendRandomReward(conn net.Conn) {
	reward := pow.GetReward()
	_, err := conn.Write([]byte(reward))
	if err != nil {
		fmt.Println("Error sending reward:", err)
	}
}
