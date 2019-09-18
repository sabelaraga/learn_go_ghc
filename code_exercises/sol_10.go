package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

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

func (m Message) print() (bool, error){
	if m.text == "" {
		return true, MyError{msg: "The message cannot be empty"}
	}
	fmt.Printf("(%d) %s\n", m.index, m.text)
	if m.isLast {
	  fmt.Println("This is supposed to be the last message")
	  return true, nil
	}
	return false, nil
}

func NewMessage(text string, index int, isLast bool) Message{
	return Message{text: text, index: index, isLast: isLast}
}

func Producer(texts [] string, msgChan chan Message) {
	for i,text := range texts {
		m := NewMessage(text, i, false)
		if i + 1 == len(msgs) {
			m.isLast = true
		}
		msgChan <- m
	}
}

func Consumer(msgChan chan Message){
       for {
		select {
		case m := <- msgChan:
		        last, err := m.print()
			if err!= nil{
				fmt.Printf("Printing failed: %t", err)
			}
			if last {
			 wait.Done()
			 return
			}
		}
	}
}

func main() {
        msgChan := make(chan Message, len(msgs))
        wait.Add(1)
        go Consumer(msgChan)
	Producer(msgs, msgChan)
	wait.Wait()
}
