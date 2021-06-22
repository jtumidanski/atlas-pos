package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type GhostGateOpen struct {
}

func (p GhostGateOpen) Name() string {
	return "ghostgate_open"
}

func (p GhostGateOpen) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "ghostgate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(990000800, 0)
		return true
	}
	script.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
