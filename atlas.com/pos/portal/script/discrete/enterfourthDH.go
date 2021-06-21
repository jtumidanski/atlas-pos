package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterFourthDH struct {
}

func (p EnterFourthDH) Name() string {
	return "enterfourthDH"
}

func (p EnterFourthDH) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(4032125) ||
		script.HasItem(l, c)(4032126) ||
		script.HasItem(l, c)(4032127) ||
		script.HasItem(l, c)(4032128) ||
		script.HasItem(l, c)(4032129) {
		script.SendPinkNotice(l, c)("ALREADY_HAVE_PROOF")
		return false
	}

	if script.QuestStarted(l, c)(20611) ||
		script.QuestStarted(l, c)(20612) ||
		script.QuestStarted(l, c)(20613) ||
		script.QuestStarted(l, c)(20614) ||
		script.QuestStarted(l, c)(20615) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 913020300) == 0 {
			_map.KillAllMonsters(l)(c.WorldId(), c.ChannelId(), 913020300)
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(913020300, 0)
			_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 913020300, 9300294, 87, 88)
			return true
		}
		script.SendPinkNotice(l, c)("SOMEONE_ALREADY_ATTEMPTING_BOSS")
		return false
	} else {
		script.SendPinkNotice(l, c)("CANNOT_ACCESS_HALL")
		return false
	}
}
