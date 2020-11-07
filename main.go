package Proxy

import (
	"net"
	"sync"
)

var rule *Rules
var once sync.Once

func NewProxyChain()*Rules{
	once.Do(
		rule = &Rules{
			tables: make([]IptableMsg,10),
			itemMap: make(map[string]Rules)
		}
	)
}


func handConn(conn net.Conn) {
	defer conn.Close()
}
