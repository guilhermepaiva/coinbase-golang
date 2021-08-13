package coinbasegolang

import (
	"net"
	"time"
)

func dialTimeout(network, addr string) (net.Conn, error) {
	var timeout = time.Duration(2 * time.Second)
	return net.DialTimeout(network, addr, timeout)
}
