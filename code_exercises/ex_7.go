// TODO: add a number to identify the message order

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

func CreateMessages(texts [] string) [] Message{
	var msgList []Message
	for _,text := range texts{
		msgList = append(msgList, NewMessage(text))
	}
	return msgList
}

func PrintList(msgList []Message){
	for _, m := range msgList{
		if err := m.print(); err!= nil{
			fmt.Printf("Printing failed: %t", err)
		}
	}
}

func main() {
	msgList := CreateMessages(msgs)
	PrintList(msgList)
}
