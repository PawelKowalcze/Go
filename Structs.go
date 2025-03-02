package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.number > 0 && mToSend.sender.name != "" && mToSend.recipient.number > 0 && mToSend.recipient.name != "" {
		return true
	}
	return false
}
func main() {
	fmt.Println(canSendMessage(messageToSend{
		message: "Hello",
		sender: user{
			name:   "John",
			number: 0,
		},
		recipient: user{
			name:   "Doe",
			number: 987654321,
		},
	}))
}
