package nets

import (
	"net"
	"time"

	"github.com/samber/lo"
	"github.com/xtaci/kcp-go/v5"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/cache"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	conns  *cache.Cache[uint32, net.Conn]
	reader Reader
}

func NewServer(reader Reader) *Server {
	return &Server{
		cache.New[uint32, net.Conn](time.Second*base.Two, time.Second),
		reader,
	}
}

func (p *Server) Run() {
	listener := lo.Must1(kcp.ListenWithOptions("0.0.0.0:"+Port, _block, _dataShards, _parityShards))

	go p.ping()

	for {
		go p.handle(lo.Must1(listener.AcceptKCP()))
	}
}

func (p *Server) Send(msg *pb.Msg) {
	data, _ := proto.Marshal(msg)

	_ = p.conns.Iterate(func(conv uint32, conn net.Conn) error {
		if conv == msg.Conv {
			return nil
		}

		if _, err := conn.Write(data); err != nil {
			logs.E.Println("remove:", conv, err)
			p.conns.Delete(conv)
		}

		return nil
	})
}

func (p *Server) Conv() uint32 { return 0 }

func (p *Server) handle(conn *kcp.UDPSession) {
	logs.D.Println("handle", conn.RemoteAddr())

	conv := conn.GetConv()
	buf := make([]byte, _bufSize)

	for {
		size, err := conn.Read(buf)
		if err != nil {
			logs.E.Println(err)
			p.conns.Delete(conv)

			return
		}

		msg := &pb.Msg{}
		_ = proto.Unmarshal(buf[:size], msg)

		if msg.Type == pb.MsgType_ping {
			p.conns.Set(conv, conn)

			continue
		}

		p.reader.Read(msg)
		p.Send(msg)
	}
}

func (p *Server) ping() {
	msg := &pb.Msg{Type: pb.MsgType_ping, Info: "ping"}

	for {
		p.Send(msg)
		time.Sleep(time.Second)
	}
}
