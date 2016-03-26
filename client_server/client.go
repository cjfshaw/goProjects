package client

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

	fmt.Fprintf(connection, "hello server")

	time.Sleep(5 * time.Second)

	fmt.Fprintf(connection, "close")

	err = connection.Close()

	if err != nil {
		fmt.Println("Error closing connection.")
		fmt.Println(err)
	}
}
