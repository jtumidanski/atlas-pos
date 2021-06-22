package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutTestWolf struct {
}

func (p OutTestWolf) Name() string {
	return "outtestWolf"
}

func (p OutTestWolf) Enter(l logrus.FieldLogger, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) != 0 {
		script.SendPinkNotice(l, c)("DEFEAT_ALL_WOLVES")
		return false
	}
	if !script.CanHold(l, c)(4001193, 1) {
		script.SendPinkNotice(l, c)("FREE_SPACE_FOR_COUSE_CLEAR_TOKEN")
		return false
	}
	script.GainItem(l, c)(4001193, 1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(140010210, 0)
	return true
}
