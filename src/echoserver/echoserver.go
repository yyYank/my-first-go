package main

import (
	"net/http"
	"net"
	"log"
	"fmt"
)

type Server interface {
	run()
	rooting(path string)
}

type EchoServer struct {
}
func (sever *EchoServer) run() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}


func hello_world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Go")
}
func goodnight_world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Good night!!!")
}
func (sever *EchoServer) rooting() {
	http.HandleFunc("/hello", hello_world)
	http.HandleFunc("/good/night", goodnight_world)
}

func main(){
	var server = &EchoServer{}
	server.rooting()
	server.run()
}
