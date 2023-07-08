package main

import (
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"github.com/xuender/poker/udps"
)

type demo struct{}

func (p *demo) Read(msg *pb.Msg) {
	logs.D.Printf("%v", msg)
}

func main() {
	room := &demo{}
	client := udps.NewClient(room)
	client.Run()

	server := udps.NewServer(room)
	server.Run()
}
