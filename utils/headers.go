package utils

import (
	"bufio"
	"strings"
)

type Headers map[string]string

func ReadHeaders(reader *bufio.Reader) (Headers, error) {
	headers := make(Headers)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return make(Headers), err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(headerParts[0])
			val := strings.TrimSpace(headerParts[1])
			headers[key] = val
		}
	}

	return headers, nil
}
