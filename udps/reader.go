package udps

import "github.com/xuender/poker/pb"

type Reader interface {
	Read(*pb.Msg)
}
