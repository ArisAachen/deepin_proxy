package Proxy

import "net"

func handConn(conn net.Conn) {
	defer conn.Close()
}
