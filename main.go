package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/Atheer-Ganayem/Go-http-from-tcp-only/utils"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Server is running on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	
	requestLine, err := utils.ReadRequestLine(reader)
	if err != nil {
		fmt.Println("Error reading request line:", err)
		return
	}

	requestHeaders, err := utils.ReadHeaders(reader)
	if err != nil {
		fmt.Println("Error reading request line:", err)
		return
	}

	req := utils.Request{RequestLine: requestLine, Headers: requestHeaders}
	err = req.ReadBody(reader)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	sl := utils.NewStatusLine(200, "ok")
	response, err := utils.ResponseJSON(sl, utils.Headers{}, utils.JM{"message": "Hello world"})
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	conn.Write([]byte(response.String()))
}