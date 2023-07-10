package nets

import "net"

type User struct {
	nick string
	addr *net.UDPAddr
}
