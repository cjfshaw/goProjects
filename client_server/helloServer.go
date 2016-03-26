package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func startListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		return listener, err
	}

	return listener, err
}

func startServer() {
	var closeConnection bool
	var message string

	listener, err := startListener()

	if err != nil {
		fmt.Println("Error occurred while trying to start the listener.")
		fmt.Println(err)
	}

	connection, err := listener.Accept()

	for closeConnection == false {
		scanner := bufio.NewScanner(connection)
		scanner.Scan()
		message = scanner.Text()
		if message != "" {
			fmt.Println("Message received:")
			fmt.Println(message)
		}

		if message == "close" {
			closeConnection = true
		}
	}

	err = listener.Close()

	if err != nil {
		fmt.Println("Error closing the connection.")
		fmt.Println(err)
	}
}

func startConnection() (net.Conn, error) {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		return connection, err
	}

	return connection, err
}

func startClient() {
	connection, err := startConnection()

	if err != nil {
		fmt.Println("Error establishing connection.")
		fmt.Println(err)
	}

	fmt.Fprintf(connection, "hello server")

	time.Sleep(5 * time.Second)

	fmt.Fprintf(connection, "close")

	err = connection.Close()

	if err != nil {
		fmt.Println("Error closing connection.")
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Starting server")
	startServer()
	fmt.Println("Starting client")
	startClient()
}
