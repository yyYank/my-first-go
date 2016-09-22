package main

import "fmt"
import "echo"

func main() {
	fmt.Println(Hello())
	echo.Echo()
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
