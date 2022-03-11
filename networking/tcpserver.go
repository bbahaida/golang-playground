package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen
	dstream, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()

	for {
		// Accept
		con, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Handle Connection
		go handle(con)
	}

}

func handle(con net.Conn) {
	for {
		data, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		if strings.TrimSpace(data) == "exit" {
			fmt.Println("close connection")
			break
		}

		fmt.Println(data)
	}
	con.Close()
}
