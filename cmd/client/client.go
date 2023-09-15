package main

import (
	"fmt"
	"net"
	"os"
	"pow_reward/internal/pow"
	"time"
)

const BufferSize = 256

func main() {
	serverURL := getServerURL()
	conn, err := establishConnection(serverURL)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", serverURL, err)
		os.Exit(1)
	}
	defer conn.Close()

	challenge, err := receiveChallenge(conn)
	if err != nil {
		fmt.Printf("Error receiving challenge: %v\n", err)
		os.Exit(1)
	}

	err = processAndSendSolution(conn, challenge)
	if err != nil {
		fmt.Printf("Error processing and sending solution: %v\n", err)
		os.Exit(1)
	}

	quote, err := receiveReward(conn)
	if err != nil {
		fmt.Printf("Error receiving reward: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Received wisdom quote: %s\n", quote)
}

func getServerURL() string {
	return os.Getenv("SERVER_URL")
}

func establishConnection(serverURL string) (net.Conn, error) {
	const timeout = 5 * time.Second // 5-second timeout
	return net.DialTimeout("tcp", serverURL, timeout)
}

func receiveChallenge(conn net.Conn) (string, error) {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error reading challenge: %w", err)
	}
	challenge := string(buffer[:n])
	fmt.Printf("Received challenge: %s\n", challenge)
	return challenge, nil
}

func processAndSendSolution(conn net.Conn, challenge string) error {
	solution := pow.Solve(challenge)

	_, err := conn.Write([]byte(solution))
	if err != nil {
		return fmt.Errorf("failed to send solution: %w", err)
	}
	fmt.Printf("Solved and sent solution: %s\n", solution)
	return nil
}

func receiveReward(conn net.Conn) (string, error) {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error getting reward: %w", err)
	}
	return string(buffer[:n]), nil
}
