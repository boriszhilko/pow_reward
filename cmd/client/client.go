package main

import (
	"fmt"
	"net"
	"os"
	"pow_reward/internal/pow"
)

const BufferSize = 256

func main() {
	serverURL := getServerURL()
	conn := establishConnection(serverURL)
	defer conn.Close()

	challenge := receiveChallenge(conn)
	processAndSendSolution(conn, challenge)

	quote := receiveReward(conn)
	fmt.Printf("Received wisdom quote: %s\n", quote)
}

func getServerURL() string {
	return os.Getenv("SERVER_URL")
}

func establishConnection(serverURL string) net.Conn {
	conn, err := net.Dial("tcp", serverURL)
	if err != nil {
		fmt.Printf("error connecting to %s: %v\n", serverURL, err)
		os.Exit(1)
	}
	return conn
}

func receiveChallenge(conn net.Conn) string {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error reading challenge:", err)
		os.Exit(1)
	}
	challenge := string(buffer[:n])
	fmt.Printf("Received challenge: %s\n", challenge)
	return challenge
}

func processAndSendSolution(conn net.Conn, challenge string) {
	solution := pow.Solve(challenge)

	_, err := conn.Write([]byte(solution))
	if err != nil {
		fmt.Println("failed to send solution:", err)
		os.Exit(1)
	}
	fmt.Printf("Solved and sent solution: %s\n", solution)
}

func receiveReward(conn net.Conn) string {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error getting reward:", err)
		os.Exit(1)
	}
	return string(buffer[:n])
}
