package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4BerserkMove struct {
}

func (p S4BerserkMove) Name() string {
	return "s4berserk_move"
}

func (p S4BerserkMove) Enter(l logrus.FieldLogger, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 0 {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(910500200, "out00")
		return true
	}
	script.SendPinkNotice(l, c)("DEFEAT_ALL_MONSTERS_FIRST")
	return true
}
