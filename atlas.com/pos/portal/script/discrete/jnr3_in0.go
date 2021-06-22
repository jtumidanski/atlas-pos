package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Jnr3In0 struct {
}

func (p Jnr3In0) Name() string {
	return "jnr3_in0"
}

func (p Jnr3In0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out1").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926110201, 0)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
