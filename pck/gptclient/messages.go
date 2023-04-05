package gptclient

import "strings"

// Parse messages string to messages string[]
// Messages format: "1. message1\n2. message2\n3. message3"
func ParseMessages(messages string) []string {
	messagesArray := strings.Split(messages, "\n")

	return messagesArray
}
