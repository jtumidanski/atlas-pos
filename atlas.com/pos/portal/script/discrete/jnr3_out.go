package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Jnr3Out struct {
}

func (p Jnr3Out) Name() string {
	return "jnr3_out"
}

func (p Jnr3Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out3").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926110203, 0)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
