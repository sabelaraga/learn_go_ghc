package main

import (
	"reflect"
	"testing"
)

type Message struct{
	text string
	index int
	isLast bool
}


func NewMessage(text string, index int, isLast bool) Message{
	return Message{text: text, index: index, isLast: isLast}
}


func Producer(texts [] string, msgChan chan Message) {
	for i,text := range texts{
		m := NewMessage(text, i, false)
		if i + 1 == len(texts) {
			m.isLast = true
		}
		msgChan <- m
	}
}

func TestProducer(t *testing.T) {
   msgChan := make(chan Message)
   msgs := []string{"!", "2"}
   go Producer(msgs, msgChan)
   var gotMsgs []string
   func (){
   	for {
		select {
		case m := <- msgChan:
		        gotMsgs = append(gotMsgs, m.text)
			if m.isLast {
			 return
			}
		}
	}
   }()
   if !reflect.DeepEqual(gotMsgs,msgs){ 
     t.Errorf("Unexpected messages produced. Got %v, want %v", gotMsgs, msgs)
   }
}

