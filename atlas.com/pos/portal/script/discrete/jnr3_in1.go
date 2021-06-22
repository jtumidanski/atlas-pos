package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Jnr3In1 struct {
}

func (p Jnr3In1) Name() string {
	return "jnr3_in1"
}

func (p Jnr3In1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out2").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926110202, 0)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
