package game

import (
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
)

type List struct {
	pokers []pb.Poker
	left   int
	top    int
	width  int
	height int
	Back   bool
}

func NewList(left, top, width, height int, pokers ...pb.Poker) *List {
	return &List{pokers: pokers, left: left, top: top, width: width, height: height}
}

func (p *List) Add(poker pb.Poker) {
	p.pokers = append(p.pokers, poker)
}

func (p *List) Click(pox, poy int) pb.Poker {
	if p.clickOut(pox, poy) {
		return pb.Poker_back
	}

	logs.D.Println("click")

	pox -= p.left

	if p.Back {
		ret := p.pokers[0]
		p.pokers = p.pokers[1:]

		return ret
	}

	pox /= 30

	if pox >= len(p.pokers) {
		pox = len(p.pokers) - 1
	}

	ret := p.pokers[pox]
	p.pokers = append(p.pokers[:pox], p.pokers[pox+1:]...)

	return ret
}

func (p *List) clickOut(pox int, poy int) bool {
	if len(p.pokers) == 0 {
		return true
	}

	if pox < p.left {
		return true
	}

	if poy < p.top {
		return true
	}

	if poy > p.top+p.height {
		return true
	}

	if pox > p.left+p.width+len(p.pokers)*30 {
		return true
	}

	if p.Back {
		if p.left > p.left+p.width+len(p.pokers) {
			return true
		}
	} else {
		if p.left > p.left+p.width+len(p.pokers)*30 {
			return true
		}
	}

	return false
}

func (p *List) Images() []*Image {
	images := make([]*Image, len(p.pokers))

	for index, poker := range p.pokers {
		img := &Image{Y: float64(p.top)}
		images[index] = img

		if p.Back {
			images[index].Poker = pb.Poker_back
			images[index].X = float64(p.left + index)
		} else {
			images[index].Poker = poker
			images[index].X = float64(p.left + index*30)
		}
	}

	return images
}
