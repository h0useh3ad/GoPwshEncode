package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"unicode/utf16"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go 'powershell command'")
		return
	}

	command := os.Args[1]

	utf16Command := utf16.Encode([]rune(command))

	utf16CommandBytes := make([]byte, len(utf16Command)*2)
	for i, v := range utf16Command {
		utf16CommandBytes[i*2] = byte(v)
		utf16CommandBytes[i*2+1] = byte(v >> 8)
	}

	encodedCommand := base64.StdEncoding.EncodeToString(utf16CommandBytes)

	fmt.Printf("\nbase64 string:\n\n")
	fmt.Printf("%v\n\n", encodedCommand)
}
