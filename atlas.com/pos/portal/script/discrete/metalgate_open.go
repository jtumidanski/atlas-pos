package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type MetalGateOpen struct {
}

func (p MetalGateOpen) Name() string {
	return "metalgate_open"
}

func (p MetalGateOpen) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "metalgate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(990000431, 0)
		return true
	}
	script.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
