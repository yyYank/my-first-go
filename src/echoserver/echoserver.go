package main



type Server interface {
	run()
	stop()
	rooting(path string)
}

type EchoServer struct {
}
func (sever *EchoServer) run() {
}
func (sever *EchoServer) stop() {
}
func (sever *EchoServer) rooting() {
}
