package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterThirdDH struct {
}

func (p EnterThirdDH) Name() string {
	return "enterthirdDH"
}

func (p EnterThirdDH) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(4032120) ||
		processor.HasItem(l, c)(4032121) ||
		processor.HasItem(l, c)(4032122) ||
		processor.HasItem(l, c)(4032123) ||
		processor.HasItem(l, c)(4032124) {
		processor.SendPinkNotice(l, c)("ALREADY_HAVE_PROOF_OF_QUALIFICATION")
		return false
	}
	if processor.QuestStarted(l, c)(20601) ||
		processor.QuestStarted(l, c)(20602) ||
		processor.QuestStarted(l, c)(20603) ||
		processor.QuestStarted(l, c)(20604) ||
		processor.QuestStarted(l, c)(20605) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 913010200) == 0 {
			_map.KillAllMonsters(l)(c.WorldId(), c.ChannelId(), 913010200)
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(913010200, 0)
			_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 913010200, 9300289, 0, 0)
			return true
		} else {
			processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_ATTEMPTING_BOSS")
			return false
		}
	}
	processor.SendPinkNotice(l, c)("LEVEL_100_SKILL_REQUIREMENT")
	return false
}
