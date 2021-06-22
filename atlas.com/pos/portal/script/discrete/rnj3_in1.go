package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Rnj3In1 struct {
}

func (p Rnj3In1) Name() string {
	return "rnj3_in1"
}

func (p Rnj3In1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj3_out2").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926100202, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
