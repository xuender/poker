package nets

import "github.com/xuender/poker/pb"

type Sender interface {
	Send(*pb.Msg)
	Conv() uint32
}
