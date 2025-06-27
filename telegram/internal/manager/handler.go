package manager

import (
	"github.com/beeper/td/bin"
	"github.com/beeper/td/mtproto"
	"github.com/beeper/td/tg"
)

// Handler abstracts updates and session handler.
type Handler interface {
	OnSession(cfg tg.Config, s mtproto.Session) error
	OnMessage(b *bin.Buffer) error
}

// NoopHandler is a noop handler.
type NoopHandler struct{}

// OnSession implements Handler.
func (n NoopHandler) OnSession(cfg tg.Config, s mtproto.Session) error {
	return nil
}

// OnMessage implements Handler
func (n NoopHandler) OnMessage(b *bin.Buffer) error {
	return nil
}
