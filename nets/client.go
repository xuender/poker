package nets

import (
	"time"

	"github.com/samber/lo"
	"github.com/xtaci/kcp-go/v5"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	conn   *kcp.UDPSession
	reader Reader
}

func NewClient(reader Reader, raddr string) *Client {
	conn := lo.Must1(kcp.DialWithOptions(raddr, _block, _dataShards, _parityShards))
	ret := &Client{conn, reader}

	go ret.ping()

	return ret
}

func (p *Client) Send(msg *pb.Msg) {
	msg.Conv = p.conn.GetConv()
	data, _ := proto.Marshal(msg)

	lo.Must1(p.conn.Write(data))
}

func (p *Client) Run() {
	buf := make([]byte, _bufSize)
	msg := &pb.Msg{}
	exit := false

	go func() {
		ticker := time.NewTicker(time.Second * base.Two)
		for range ticker.C {
			if exit {
				logs.E.Println("exit")
				ticker.Stop()
				p.conn.Close()
			}

			exit = true
		}
	}()

	for {
		size := lo.Must1(p.conn.Read(buf))
		_ = proto.Unmarshal(buf[:size], msg)

		if msg.Type == pb.MsgType_ping {
			exit = false

			continue
		}

		p.reader.Read(msg)
	}
}

func (p *Client) Conv() uint32 { return p.conn.GetConv() }
func (p *Client) ping() {
	msg := &pb.Msg{Type: pb.MsgType_ping, Info: "ping"}

	for {
		p.Send(msg)
		time.Sleep(time.Second)
	}
}
