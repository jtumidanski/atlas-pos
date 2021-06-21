package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterPort struct {
}

func (p EnterPort) Name() string {
	return "enterPort"
}

func (p EnterPort) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(21301) &&
		script.QuestProgressInt(l, c)(21301, 9001013) == 0 {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010700) != 0 {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_CHALLENGING_THIEF_CROW")
			return false
		}
		_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 108010700, 9001013, 2732, 3)
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(108010700, "west00")
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(140020300, 1)
	return true
}
