package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello"
}

type Event struct {}
type Handler interface {
	handle(Event)
}

type HogeHandler struct{
}

func handle(e Event) {
	 //do something
}
