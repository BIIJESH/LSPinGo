package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d \r\n\r\n%s", len(content), content)
}
func DecodeMessage (msg []byte){
header , content , found := bytes.Cut(msg)
}
