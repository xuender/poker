package game

import (
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/los"
	"github.com/xuender/poker/nets"
	"github.com/xuender/poker/pb"
)

type Script struct {
	bus      *Bus
	users    []string
	isServer bool
	sender   nets.Sender
}

func NewScript(bus *Bus) *Script {
	ret := &Script{bus: bus, users: []string{}}
	bus.SetReader(ret)

	return ret
}

// nolint: cyclop
func (p *Script) Read(msg *pb.Msg) {
	switch msg.Type {
	case pb.MsgType_ping:
		// nolint: gocritic, exhaustive
		switch p.bus.Scene() {
		case pb.Scene_start:
			p.users = msg.Users

			if p.isServer {
				// nolint: gocritic
				p.bus.Start = append(p.users, "等待加入...")
				if len(p.users) > 1 {
					logs.D.Println(">>>>>>>", "run")

					msg := &pb.Msg{Type: pb.MsgType_run}
					p.sender.Send(msg)

					go func() {
						time.Sleep(base.Three * time.Second)
						p.Read(msg)
					}()
				}
			} else {
				// nolint: gocritic
				p.bus.Start = append(p.users, "等待开始...")
			}
		}
	case pb.MsgType_reset:
		p.bus.Backs = msg.Backs
	case pb.MsgType_run:
		logs.W.Println("游戏开始")
		p.bus.To(pb.Scene_desktop)

		if p.isServer {
			// 重置
			msg := &pb.Msg{Type: pb.MsgType_reset}
			pokers := make([]pb.Poker, _pokerNum)

			for i := 1; i <= 54; i++ {
				pokers[i-1] = pb.Poker(i)
			}

			pokers = lo.Shuffle(pokers)
			msg.Backs = pokers

			p.Read(msg)
			p.sender.Send(msg)
		}
	case pb.MsgType_take:
		if p.sender.Conv() == msg.Conv {
			p.bus.My = append(p.bus.My, msg.Take)
		} else {
			p.bus.Your = append(p.bus.Your, msg.Take)
		}

		los.Pull(p.bus.Backs, msg.Take)
	}
}

func (p *Script) Take(poker pb.Poker) {
	p.sender.Send(&pb.Msg{Type: pb.MsgType_take, Take: poker})
}
