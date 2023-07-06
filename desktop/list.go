package desktop

import (
	"image"

	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/pb"
)

type List struct {
	pokers []pb.Poker
	bounds image.Rectangle
	x      int
	y      int
	Back   bool
}

func NewList(x, y int, bounds image.Rectangle, pokers ...pb.Poker) *List {
	return &List{pokers: pokers, x: x, y: y, bounds: bounds}
}

func (p *List) Add(poker pb.Poker) {
	p.pokers = append(p.pokers, poker)
}

func (p *List) Click(pox, poy int) pb.Poker {
	if p.clickOut(pox, poy) {
		return pb.Poker_back
	}

	logs.D.Println("click")

	pox -= p.x

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

	if pox < p.x {
		return true
	}

	if poy < p.y {
		return true
	}

	if poy > p.y+p.bounds.Dy() {
		return true
	}

	if pox > p.x+p.bounds.Dx()+len(p.pokers)*30 {
		return true
	}

	if p.Back {
		if p.x > p.x+p.bounds.Dx()+len(p.pokers) {
			return true
		}
	} else {
		if p.x > p.x+p.bounds.Dx()+len(p.pokers)*30 {
			return true
		}
	}

	return false
}

func (p *List) Images() []*Image {
	images := make([]*Image, len(p.pokers))

	for index, poker := range p.pokers {
		img := &Image{Y: float64(p.y)}
		images[index] = img

		if p.Back {
			images[index].Poker = pb.Poker_back
			images[index].X = float64(p.x + index)
		} else {
			images[index].Poker = poker
			images[index].X = float64(p.x + index*30)
		}
	}

	return images
}
