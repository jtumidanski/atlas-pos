package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterDisguise1 struct {
}

func (p EnterDisguise1) Name() string {
	return "enterDisguise1"
}

func (p EnterDisguise1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(20301) ||
		script.QuestStarted(l, c)(20302) ||
		script.QuestStarted(l, c)(20303) ||
		script.QuestStarted(l, c)(20304) ||
		script.QuestStarted(l, c)(20305) {
		if len(_map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), 108010600)) > 0 {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if script.HasItem(l, c)(4032101) {
			script.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(108010600, "out00")
		return true
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(130010010, "out00")
	return true
}
