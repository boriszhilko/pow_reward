package main

import (
	"fmt"
	"net"
	"os"
	"pow_reward/internal/pow"
)

const BufferSize = 256

func main() {
	listener, err := createServer(":8080")
	if err != nil {
		fmt.Println("Error creating server:", err)
		os.Exit(1)
	}
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

func createServer(address string) (net.Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("Error listening on %s: %v", address, err)
	}
	fmt.Printf("Listening on %s\n", address)
	return listener, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	if err := interactWithClient(conn); err != nil {
		fmt.Println("Error during client interaction:", err)
	}
}

func interactWithClient(conn net.Conn) error {
	challenge, err := sendChallenge(conn)
	if err != nil {
		return err
	}

	response, err := receiveResponse(conn)
	if err != nil {
		return err
	}

	isCorrect := pow.VerifySolution(challenge, response)
	return sendFeedback(conn, isCorrect)
}

func sendChallenge(conn net.Conn) (string, error) {
	challenge := pow.GetChallenge(10)
	_, err := conn.Write([]byte(challenge))
	if err != nil {
		return "", err
	}
	return challenge, nil
}

func receiveResponse(conn net.Conn) (string, error) {
	buffer := make([]byte, BufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func sendFeedback(conn net.Conn, isCorrect bool) error {
	if isCorrect {
		return sendRandomReward(conn)
	}
	_, err := conn.Write([]byte("Incorrect PoW"))
	return err
}

func sendRandomReward(conn net.Conn) error {
	reward := pow.GetReward()
	_, err := conn.Write([]byte(reward))
	return err
}
