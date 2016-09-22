package main

import "fmt"

type Event interface {
	getEventName() string
}

type HogeEvent struct {
}
func(e *HogeEvent) getEventName() string {
	return "hoge"
}

type Handler interface {
	handle(Event)
}

type HogeHandler struct {
}

func (h *HogeHandler) handle(e Event) {
	switch e.getEventName(){
	case "onClick": fmt.Println("a")
	case "onClose": fmt.Println("b")
	case "onDoubleClick": fmt.Println("c")
	default:fmt.Println("unexpected event")
	}
}


func main(){
	var handler = &HogeHandler{}
	handler.handle(&HogeEvent{})
}
