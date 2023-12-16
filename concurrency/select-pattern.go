package main

import (
	"fmt"
)

type Server struct {
	quitch chan bool // 종료 flag
	msgCh  chan string
}

func NewServer() *Server {
	return &Server{
		quitch: make(chan bool, 1),
		msgCh:  make(chan string, 128),
	}
}

func (s *Server) Start() {
	fmt.Println("server start")
	s.loop() // block
}

func (s *Server) loop() {
	// for-select pattern을 사용해서 -> channel을 구현할 수 있음
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Println(msg)

		// 종료여부 검사
		case <-s.quitch:
			// do some stuff we need to quit
			break

		default:
			// continue
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we recevied a message ", msg)
	s.msgCh <- msg
}

func (s *Server) handleExit() {
	fmt.Println("We Ready Program stop")
	s.quitch <- true
}

func main() {
	s := NewServer()

	// must attach "go" keward
	go s.Start()

	for i := 0; i < 100; i++ {
		s.handleMessage(string(i))
	}
	s.handleExit()
}
