package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	serverIP := os.Getenv("MARKET_SERVER_IP")
	if serverIP == "" {
		serverIP = "127.0.0.1"
	}

	conn, err := net.Dial("tcp", serverIP + ":5005")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server.")
	buf := make([]byte, 48)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Disconnected or error:", err)
			break
		}

		if n == 0 {
			continue
		}

		fmt.Printf("Received %d bytes: % x\n", n, buf[:n])
	}
}