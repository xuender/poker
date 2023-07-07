package udps

import (
	"math/rand"
	"net"
	"os/user"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"google.golang.org/protobuf/proto"
)

const _port = 3880

type Server struct {
	users      map[int64]*User
	nick       string
	conn       *net.UDPConn
	serverAddr *net.UDPAddr
	id         int64
	reader     Reader
}

func NewServer(reader Reader) *Server {
	rand.Seed(time.Now().UnixNano())

	user := lo.Must1(user.Current())
	ret := &Server{
		users:  map[int64]*User{},
		nick:   user.Username,
		reader: reader,
		id:     rand.Int63(),
	}

	return ret
}

func (p *Server) Run() {
	p.client()
}

func (p *Server) SendByID(msg *pb.Msg, client int64) {
	msg.Client = client
	data := lo.Must1(proto.Marshal(msg))

	if p.serverAddr == nil {
		for id, user := range p.users {
			_, err := p.conn.WriteToUDP(data, user.addr)
			if err != nil {
				delete(p.users, id)
				p.Send(msg)
			}
		}
	} else {
		_, err := p.conn.WriteToUDP(data, p.serverAddr)
		if err != nil {
			logs.E.Println(err)

			return
		}
	}
}

func (p *Server) Send(msg *pb.Msg) {
	logs.D.Println("send", msg)
	p.SendByID(msg, p.id)
}

func (p *Server) read(data []byte) {
	msg := &pb.Msg{}

	lo.Must0(proto.Unmarshal(data, msg))
	p.reader.Read(msg)
}

func (p *Server) AddAddr(client int64, user *User) {
	p.users[client] = user
	logs.D.Printf("client: %d: %v", client, user)
}

func (p *Server) client() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: _port,
	})
	if err != nil {
		logs.E.Println("err:", err)

		return
	}
	defer conn.Close()

	msg := &pb.Msg{Nick: p.nick + "_c", Client: p.id}
	sendData, _ := proto.Marshal(msg)

	_, err = conn.Write(sendData)
	if err != nil {
		logs.E.Println("发送数据失败", err)
		return
	}

	p.conn = conn
	data := make([]byte, 4096)

	for {
		udp, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			logs.W.Println("接受数据失败", err)
			p.server()

			return
		}

		p.serverAddr = addr
		p.read(data[:udp])
	}
}

func (p *Server) server() {
	conn := lo.Must1(net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: _port,
	}))
	defer conn.Close()
	p.conn = conn

	for {
		var bf [1024]byte

		size, addr := lo.Must2(conn.ReadFromUDP(bf[:]))
		msg := &pb.Msg{}

		lo.Must0(proto.Unmarshal(bf[:size], msg))
		logs.D.Printf("data:%v, addr:%v", msg, addr)

		p.AddAddr(msg.Client, &User{nick: msg.Nick, addr: addr})

		nicks := make([]string, 0, len(p.users)+1)
		for _, user := range p.users {
			nicks = append(nicks, user.nick)
		}

		nicks = append(nicks, p.nick)

		msg.Users = nicks
		p.reader.Read(msg)
		p.SendByID(msg, msg.Client)
	}
}
