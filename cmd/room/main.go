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
	server := udps.NewServer(&demo{})
	server.Run()
}
