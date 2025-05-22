// utils/encoder.go
package utils

import "fmt"

// TextToBinary converts a message (string) to a binary string (e.g., "hi" -> "0110100001101001")
func TextToBinary(message string) string {
	binary := ""
	for i := 0; i < len(message); i++ {
		binary += fmt.Sprintf("%08b", message[i]) // convert each char to 8-bit binary
	}
	return binary
}
