package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	CRLF      = "\r\n"
	SEPARATOR = " "
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer c.Close()

	handleConn(c)
}

func handleConn(c net.Conn) {
	buf := make([]byte, 1024)
	if _, err := c.Read(buf); err != nil {
		fmt.Println("Failed to read from connection")
		os.Exit(1)
	}

	req := string(buf)

	lines := strings.Split(req, CRLF)

	path := strings.Split(lines[0], SEPARATOR)[1]

	var response string

	if path == "/" {
		response = fmt.Sprintf("HTTP/1.1 200 OK%s%s", CRLF, CRLF)
	} else {
		response = fmt.Sprintf("HTTP/1.1 404 Not Found%s%s", CRLF, CRLF)
	}

	c.Write([]byte(response))
}
