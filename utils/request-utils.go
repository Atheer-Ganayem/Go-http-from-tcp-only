package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	RequestLine
	Headers
	Body string
}

type RequestLine struct {
	Method  string
	Path    string
	Protocol string
}


func ReadRequestLine(reader *bufio.Reader) (RequestLine, error) {
	return RequestLine{}, errors.New("test")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return RequestLine{}, err
	}

	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")

	if len(parts) != 3 {
		return RequestLine{}, errors.New("incorrect request line")
	}

	return RequestLine{parts[0], parts[1], parts[2]}, nil
}


func (req *Request) ReadBody(reader *bufio.Reader) (error) {
	if cl, ok := req.Headers["Content-Length"]; ok {
		var length int

		_, err := fmt.Sscanf(cl, "%d", &length)
		if err != nil || length <= 0 {
			return errors.New("Invalid Content-Length")
		}

		bodyBytes := make([]byte, length)
		_, err = reader.Read(bodyBytes)
		if err != nil {
			return errors.New("Couldn't read body")
		}

		req.Body = string(bodyBytes)
	}

	return nil
}

func (req *Request) Print() {
	fmt.Println(req.Method, req.Path, req.Protocol)
	
	for key, val := range req.Headers {
		fmt.Println(key, val)
	}

	fmt.Println(req.Body)

}

func (req *Request) BodyJSON(target *interface{}) error {
	if req.Headers["Content-Type"] != "application/json" {
		return errors.New("invalid content type: expected application/json")
	}

	return json.Unmarshal([]byte(req.Body), target)
}