package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4FireHawk struct {
}

func (p S4FireHawk) Name() string {
	return "s4firehawk"
}

func (p S4FireHawk) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(6240) {
		script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 921100200) != 0 {
		script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 921100200)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(921100200, 0)
	return true
}
