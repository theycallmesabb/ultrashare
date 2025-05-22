package utils

import "fmt"

func Texttobinary(message string) string {
	binary := ""
	for i := 0; i < len(message); i++ {
		binary += fmt.Sprintf("%08b", message[i])
	}
	return binary
}
