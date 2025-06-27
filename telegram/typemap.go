package telegram

import (
	"sync"

	"github.com/beeper/td/mt"
	"github.com/beeper/td/proto"
	"github.com/beeper/td/tg"
	"github.com/beeper/td/tmap"
)

// Port is default port used by telegram.
const Port = 443

var (
	typesMap  *tmap.Map
	typesOnce sync.Once
)

func getTypesMapping() *tmap.Map {
	typesOnce.Do(func() {
		typesMap = tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		)
	})
	return typesMap
}
