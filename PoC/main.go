package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating server:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Go: Waiting for connection...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Go: Received %s\n", message)

		if message == "ping" {
			response := "pong\n"
			fmt.Printf("Go: Sending %s\n", response)
			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error writing:", err.Error())
				return
			}
		} else {
			fmt.Println("Go: Unexpected message")
			return
		}
	}
}
