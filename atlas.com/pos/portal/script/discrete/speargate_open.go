package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type SpearGateOpen struct {
}

func (p SpearGateOpen) Name() string {
	return "speargate_open"
}

func (p SpearGateOpen) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "speargate").State() == 4 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(990000401, 0)
		return true
	}
	script.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
