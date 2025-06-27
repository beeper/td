package pool

import (
	"context"

	"github.com/go-faster/errors"
	"go.uber.org/atomic"

	"github.com/beeper/td/bin"
	"github.com/beeper/td/tdsync"
)

// ErrConnDead means that connection is closed and can't be used anymore.
var ErrConnDead = errors.New("connection dead")

// Conn represents Telegram MTProto connection.
type Conn interface {
	Run(ctx context.Context) error
	Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error
	Ping(ctx context.Context) error
	Ready() <-chan struct{}
}

type poolConn struct {
	Conn
	id      int64 // immutable
	dc      *DC   // immutable
	deleted *atomic.Bool
	dead    *tdsync.Ready
}

func (p *poolConn) Dead() <-chan struct{} {
	return p.dead.Ready()
}
