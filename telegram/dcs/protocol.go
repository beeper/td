package dcs

import (
	"net"

	"github.com/beeper/td/transport"
)

type protocol interface {
	Handshake(conn net.Conn) (transport.Conn, error)
}
