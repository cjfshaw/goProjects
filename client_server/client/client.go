package main

import (
	"fmt"
	"net"
	"time"
)

func startClient() (net.Conn, error) {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		return connection, err
	}

	return connection, err
}

func main() {
	connection, err := startClient()

	if err != nil {
		fmt.Println("Error establishing connection.")
		fmt.Println(err)
	}

	fmt.Fprintf(connection, "hello server\n")
	fmt.Println("First message 'hello server' sent.")

	fmt.Println("Sleeping for 5 seconds.")
	time.Sleep(5 * time.Second)

	fmt.Fprintf(connection, "close\n")
	fmt.Println("Second message 'close' sent.")

	err = connection.Close()

	if err != nil {
		fmt.Println("Error closing connection.")
		fmt.Println(err)
	} else {
		fmt.Println("Connection closed, client shutting down!")
	}
}
