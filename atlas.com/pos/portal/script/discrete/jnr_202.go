package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Jnr202 struct {
}

func (p Jnr202) Name() string {
	return "jnr_202"
}

func (p Jnr202) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr32_out").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926110200, 2)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
