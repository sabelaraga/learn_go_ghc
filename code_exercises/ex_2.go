// TODO: add a function that prints the message.

package main

import (
	"fmt"
)

type Message struct{
	text string
}


func main() {
	m := Message{text:"Welcome to GHC!"}
	print(m)
}
