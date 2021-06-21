package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterThirdDH struct {
}

func (p EnterThirdDH) Name() string {
	return "enterthirdDH"
}

func (p EnterThirdDH) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(4032120) ||
		script.HasItem(l, c)(4032121) ||
		script.HasItem(l, c)(4032122) ||
		script.HasItem(l, c)(4032123) ||
		script.HasItem(l, c)(4032124) {
		script.SendPinkNotice(l, c)("ALREADY_HAVE_PROOF_OF_QUALIFICATION")
		return false
	}
	if script.QuestStarted(l, c)(20601) ||
		script.QuestStarted(l, c)(20602) ||
		script.QuestStarted(l, c)(20603) ||
		script.QuestStarted(l, c)(20604) ||
		script.QuestStarted(l, c)(20605) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 913010200) == 0 {
			_map.KillAllMonsters(l)(c.WorldId(), c.ChannelId(), 913010200)
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(913010200, 0)
			_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 913010200, 9300289, 0, 0)
			return true
		} else {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_ATTEMPTING_BOSS")
			return false
		}
	}
	script.SendPinkNotice(l, c)("LEVEL_100_SKILL_REQUIREMENT")
	return false
}
