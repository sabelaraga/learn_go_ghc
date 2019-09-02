// TODO: validate that the message is not empty when printing

package main

import (
	"fmt"
)

type MyError struct{
	msg string
}

func (e MyError) Error() string{
	return e.msg
}

type Message struct{
	text string
}

func (m Message) print() error{

	fmt.Println(m.text)
	return nil
}

func NewMessage(text string) Message{
	return Message{text: text}
}

func main() {
	m := NewMessage("Welcome to GHC!")
	m.print()
}
