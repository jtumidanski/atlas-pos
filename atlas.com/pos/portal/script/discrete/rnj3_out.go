package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Rnj3Out struct {
}

func (p Rnj3Out) Name() string {
	return "rnj3_out"
}

func (p Rnj3Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj3_out3").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926100203, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
