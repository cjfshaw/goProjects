package server

import (
	"bufio"
	"fmt"
	"net"
)

func startListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		return listener, err
	}

	return listener, err
}

func main() {
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
