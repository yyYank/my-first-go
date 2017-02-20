package main

import (
	"net/http"
)

type Server interface {
	run()
	stop()
	rooting(path string)
}


func handle(w http.ResponseWriter, r *http.Request) {

}
type EchoServer struct {
}
func (sever *EchoServer) run() {

	http.ListenAndServe(":8080", nil)
}
func (sever *EchoServer) stop() {
}
func (sever *EchoServer) rooting() {
	http.HandleFunc("/", handle)
}
