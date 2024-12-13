package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

// func DecodeMessage(msg []byte) (int, error) {
// 	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
// 	if !found {
// 		return 0, errors.New("Did not find seperator")
// 	}
// 	contentLengthBytes := header[len("Content-Length: "):]
// 	contentLength, err := strconv.Atoi(string(contentLengthBytes))
// 	if err != nil {
// 		return 0, err
// 	}
// 	//TODO:do something of this content
// 	_ = content
// 	return contentLength, nil
// }

func DecodeMessage(msg []byte) (int, error) {
	// Split the message into header and content
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("did not find separator")
	}

	// Ensure the header starts with "Content-Length: "
	headerStr := string(header)
	const prefix = "Content-Length: "
	if !bytes.HasPrefix(header, []byte(prefix)) {
		return 0, errors.New("invalid header: missing Content-Length")
	}

	// Extract and clean the Content-Length value
	contentLengthStr := headerStr[len(prefix):]                          // Extract everything after "Content-Length: "
	contentLengthStr = string(bytes.TrimSpace([]byte(contentLengthStr))) // Trim extra spaces

	// Convert to integer
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return 0, fmt.Errorf("invalid content length '%s': %w", contentLengthStr, err)
	}

	// Debug the content (optional)
	_ = content

	return contentLength, nil
}
