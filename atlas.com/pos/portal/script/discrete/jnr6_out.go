package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Jnr6Out struct {
}

func (p Jnr6Out) Name() string {
	return "jnr6_out"
}

func (p Jnr6Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr6_out").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926110300, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
