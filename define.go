package Proxy

import "log"

const (
	TCP  = "tcp"
	HTTP = "http"
)

type TableTarget int

func (t TableTarget) ToString() string {
	var target string
	switch t {
	case ACCEPT:
		target = "ACCEPT"
	case DROP:
		target = "DROP"
	case REJECT:
		target = "REJECT"
	case SNAT:
		target = "SNAT"
	case MASQUERADE:
		target = "MASQUERADE"
	case DNAT:
		target = "DNAT"
	case REDIRECT:
		target = "REDIRECT"
	case LOG:
		target = "LOG"
	default:
		log.Fatal("ip tables target is not exist")
	}
	return target
}

const (
	ACCEPT TableTarget = iota
	DROP
	REJECT
	SNAT
	MASQUERADE
	DNAT
	REDIRECT
	LOG
)
