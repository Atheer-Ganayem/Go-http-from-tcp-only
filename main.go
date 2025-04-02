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
		errorRes(conn, "Error reading request line: " + err.Error())
		return
	}

	requestHeaders, err := utils.ReadHeaders(reader)
	if err != nil {
		errorRes(conn, "Error reading headers: " + err.Error())
		return
	}

	req := utils.Request{RequestLine: requestLine, Headers: requestHeaders}
	err = req.ReadBody(reader)
	if err != nil {
		errorRes(conn, "Error reading body: " + err.Error())
		return
	}

	sl := utils.NewStatusLine(200, "ok")
	response, err := utils.ResponseJSON(sl, utils.Headers{}, utils.JM{"message": "Hello world"})
	if err != nil {
		errorRes(conn, "Error parsing response: " + err.Error())
		return
	}

	conn.Write([]byte(response.String()))
}

func errorRes(conn net.Conn, msg string) {
	sl := utils.NewStatusLine(400, "Bad Request")
	response, err := utils.ResponseJSON(sl, utils.Headers{}, utils.JM{"message": msg})
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	conn.Write([]byte(response.String()))
}