// TODO: add a boolean to signal the last message

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
	index int
	isLast bool
}

func (m Message) print() error{
	if m.text == "" {
		return MyError{msg: "The message cannot be empty"}
	}
	fmt.Printf("(%d) %s\n", m.index, m.text)
	if m.isLast {
	  fmt.Println("This is supposed to be the last message")
	}
	return nil
}

func NewMessage(text string, index int, isLast bool) Message{
	return Message{text: text, index: index, isLast: isLast}
}

func CreateMessages(texts [] string) [] Message{
	var msgList []Message
	for i,text := range texts{
		msgList = append(msgList, NewMessage(text, i, false))
	}
	msgList[len(msgList)-1].isLast = true
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
