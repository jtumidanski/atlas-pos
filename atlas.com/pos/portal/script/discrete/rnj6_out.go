package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Rnj6Out struct {
}

func (p Rnj6Out) Name() string {
	return "rnj6_out"
}

func (p Rnj6Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj6_out").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926100300, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}