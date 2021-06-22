package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type StatueGateOpen struct {
}

func (p StatueGateOpen) Name() string {
	return "statuegate_open"
}

func (p StatueGateOpen) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "statuegate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(990000301, 0)
		return true
	}
	script.SendPinkNotice(l, c)("GATE_IS_CLOSED")
	return false
}
