package nets

import "github.com/xuender/poker/pb"

type Reader interface {
	Read(*pb.Msg)
}
