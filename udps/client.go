package udps

import (
	"math/rand"
	"net"
	"os/user"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	nick       string
	id         int64
	conn       *net.UDPConn
	serverAddr *net.UDPAddr
	reader     Reader
}

// nolint: gosec
func NewClient(reader Reader) *Client {
	rand.Seed(time.Now().UnixNano())

	user := lo.Must1(user.Current())

	return &Client{id: rand.Int63(), nick: user.Name, reader: reader}
}

func (p *Client) Run() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: _port,
	})
	if err != nil {
		logs.E.Println("err:", err)

		return
	}
	defer conn.Close()

	p.ping(conn, true)
	go p.ping(conn, false)

	p.conn = conn
	data := make([]byte, base.Kilo*base.Four)

	for {
		udp, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			logs.W.Println("接受数据失败", err)
			p.conn.Close()

			return
		}

		p.serverAddr = addr
		p.read(data[:udp])
	}
}

func (p *Client) ping(conn *net.UDPConn, one bool) {
	msg := &pb.Msg{Nick: p.nick + "_c", Client: p.id, Type: pb.MsgType_ping}
	sendData, _ := proto.Marshal(msg)

	for {
		if !one {
			time.Sleep(time.Second)
		}

		_, err := conn.Write(sendData)
		if err != nil || one {
			return
		}
	}
}

func (p *Client) read(data []byte) {
	msg := &pb.Msg{}

	lo.Must0(proto.Unmarshal(data, msg))
	p.reader.Read(msg)
}

func (p *Client) Send(msg *pb.Msg) {
	_, err := p.conn.WriteToUDP(lo.Must1(proto.Marshal(msg)), p.serverAddr)
	logs.Log(err)
}
