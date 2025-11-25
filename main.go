package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5005")
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