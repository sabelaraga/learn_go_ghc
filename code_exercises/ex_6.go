// TODO: create and print a list of messages

package main

import (
	"fmt"
)

var (
  msgs = []string {"Welcome to GHC!", "This", "is", "a super great", "conference!"}
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
	if m.text == "" {
		return MyError{msg: "The message cannot be empty"}
	}
	fmt.Println(m.text)
	return nil
}

func NewMessage(text string) Message{
	return Message{text: text}
}

func main() {
	m := NewMessage(msgs[0])
	if err := m.print(); err!= nil{
		fmt.Printf("Printing failed: %t", err)
	}
}
