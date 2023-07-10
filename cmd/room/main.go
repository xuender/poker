package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/nets"
	"github.com/xuender/poker/pb"
)

type demo struct{}

func (p *demo) Read(msg *pb.Msg) {
	logs.W.Printf("%s: %s\n", msg.Nick, msg.Info)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	room := &demo{}

	var sender nets.Sender

	if flag.NArg() > 0 {
		client := nets.NewClient(room, flag.Arg(0))
		sender = client

		go client.Run()
	} else {
		server := nets.NewServer(room)
		sender = server

		go server.Run()
	}

	logs.I.Println("请输入昵称:")

	inputReader := bufio.NewReader(os.Stdin)
	nick := lo.Must1(inputReader.ReadString('\n'))
	nick = nick[:len(nick)-1]

	logs.I.Println("你好,", nick)

	for {
		info := lo.Must1(inputReader.ReadString('\n'))
		info = info[:len(info)-1]
		sender.Send(&pb.Msg{Nick: nick, Info: info, Type: pb.MsgType_take})
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "room\n\n")
	fmt.Fprintf(os.Stderr, "聊天室.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [ipaddress]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
