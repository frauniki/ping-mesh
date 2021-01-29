package domain

import (
	"net"
)

type IPAddr struct {
	IP   net.IP `json:"ip"`
	Zone string `json:"zone"`
}
