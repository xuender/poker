package udps

import (
	"math/rand"
	"net"
	"os/user"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/cache"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
	"google.golang.org/protobuf/proto"
)

const _port = 3880

type Server struct {
	users  *cache.Cache[int64, *User]
	nick   string
	conn   *net.UDPConn
	id     int64
	reader Reader
}

// nolint: gosec
func NewServer(reader Reader) *Server {
	rand.Seed(time.Now().UnixNano())

	user := lo.Must1(user.Current())
	ret := &Server{
		users:  cache.New[int64, *User](time.Second*base.Two, time.Second),
		nick:   user.Username,
		reader: reader,
		id:     rand.Int63(),
	}

	return ret
}

func (p *Server) SendByID(msg *pb.Msg, client int64) {
	msg.Client = client
	data := lo.Must1(proto.Marshal(msg))

	_ = p.users.Iterate(func(key int64, user *User) error {
		_, err := p.conn.WriteToUDP(data, user.addr)
		if err != nil {
			p.users.Delete(key)
			p.Send(msg)
		}

		return nil
	})
}

func (p *Server) Send(msg *pb.Msg) {
	logs.D.Println("send", msg)
	p.SendByID(msg, p.id)
}

func (p *Server) AddAddr(client int64, user *User) {
	p.users.Set(client, user)
	logs.D.Printf("client: %d: %v", client, user)
}

func (p *Server) Run() {
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

		if msg.Type == pb.MsgType_ping {
			p.AddAddr(msg.Client, &User{nick: msg.Nick, addr: addr})

			continue
		}

		nicks := make([]string, 0, p.users.Len()+1)
		_ = p.users.Iterate(func(_ int64, user *User) error {
			nicks = append(nicks, user.nick)

			return nil
		})

		nicks = append(nicks, p.nick)

		msg.Users = nicks
		p.reader.Read(msg)
		p.SendByID(msg, msg.Client)
	}
}
