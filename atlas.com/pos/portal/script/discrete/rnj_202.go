package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type Rnj202 struct {
}

func (p Rnj202) Name() string {
	return "rnj_202"
}

func (p Rnj202) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj32_out").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(926100200, 2)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
