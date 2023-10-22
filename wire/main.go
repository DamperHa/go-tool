package main

import (
	"fmt"
)

type Message string

func NewMessage() Message { return Message("Hi there") }

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()

	fmt.Println(msg)
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

//func main.go() {
//	e := NewEvent(NewGreeter(NewMessage()))
//	e.Start()
//}

func main() {
	e := InitializeEvent()
	e.Start()
}
