package game

import (
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
)

type List struct {
	left   int
	top    int
	width  int
	height int
	Back   bool
}

func NewList(left, top, width, height int) *List {
	return &List{left: left, top: top, width: width, height: height}
}

func (p *List) Click(pox, poy int, pokers []pb.Poker) pb.Poker {
	if p.clickOut(pox, poy, pokers) {
		return pb.Poker_back
	}

	logs.D.Println("click")

	pox -= p.left

	if p.Back {
		ret := pokers[0]
		pokers = pokers[1:]

		return ret
	}

	pox /= 30

	if pox >= len(pokers) {
		pox = len(pokers) - 1
	}

	ret := pokers[pox]
	pokers = append(pokers[:pox], pokers[pox+1:]...)

	return ret
}

func (p *List) clickOut(pox int, poy int, pokers []pb.Poker) bool {
	if len(pokers) == 0 {
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

	if pox > p.left+p.width+len(pokers)*30 {
		return true
	}

	if p.Back {
		if p.left > p.left+p.width+len(pokers) {
			return true
		}
	} else {
		if p.left > p.left+p.width+len(pokers)*30 {
			return true
		}
	}

	return false
}

func (p *List) Images(pokers []pb.Poker) []*Image {
	images := make([]*Image, len(pokers))

	for index, poker := range pokers {
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
