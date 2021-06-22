package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Rnj3In0 struct {
}

func (p Rnj3In0) Name() string {
	return "rnj3_in0"
}

func (p Rnj3In0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj3_out1").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926100201, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
