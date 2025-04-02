package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Response struct {
	StatusLine
	Headers
	Body string
}

type StatusLine struct {
	Version string
	Code    int
	Message string
}

type JM map[string]string

func NewStatusLine(code int, msg string) (StatusLine) {
	return StatusLine{"HTTP/1.1", code, msg}
}

func ResponseJSON(sl StatusLine, headers Headers, body map[string]string) (Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return Response{}, errors.New("Couldn't convert body to json")
	}

	headers["Content-Type"] = "application/json"
	headers["Content-Length"] = strconv.Itoa(len(jsonBody))

	return Response{StatusLine: sl, Headers: headers, Body: string(jsonBody)}, nil
}

func (res *Response) String() string {
	responseString := fmt.Sprintf("%s %d %s\r\n", res.Version, res.Code, res.Message)

	for key, val := range res.Headers {
		responseString += fmt.Sprintf("%s: %s\r\n", key, val)
	}

	responseString += "\r\n"
	responseString += res.Body

	return responseString
}