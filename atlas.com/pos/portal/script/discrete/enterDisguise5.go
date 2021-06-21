package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterDisguise5 struct {
}

func (p EnterDisguise5) Name() string {
	return "enterDisguise5"
}

func (p EnterDisguise5) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(20301) ||
		script.QuestStarted(l, c)(20302) ||
		script.QuestStarted(l, c)(20303) ||
		script.QuestStarted(l, c)(20304) ||
		script.QuestStarted(l, c)(20305) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010640) > 0 {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if script.HasItem(l, c)(4032105) {
			script.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(108010640, "east00")
		return true
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(130020000, "east00")
	return true
}
