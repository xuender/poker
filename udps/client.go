package udps

import (
	"net"
	"time"

	"github.com/lithdew/reliable"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

type Client struct {
	endpoint *reliable.Endpoint
}

func NewClient() {
}

func (p *Client) handler(buf []byte, _ net.Addr) {
	logs.D.Println(buf)
}

func (p *Client) Run() {
	add := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: _port,
	}
	conn := lo.Must1(net.ListenPacket("udp", "0.0.0.0:3831"))
	p.endpoint = reliable.NewEndpoint(conn, reliable.WithEndpointPacketHandler(p.handler))
	go p.endpoint.Listen()

	for {
		// TODO data
		if err := p.endpoint.WriteReliablePacket([]byte{}, add); err != nil {
			return
		}

		time.Sleep(time.Second)
	}
}
