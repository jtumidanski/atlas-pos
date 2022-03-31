package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterFourthDH struct {
}

func (p EnterFourthDH) Name() string {
	return "enterfourthDH"
}

func (p EnterFourthDH) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(4032125) ||
		processor.HasItem(l, c)(4032126) ||
		processor.HasItem(l, c)(4032127) ||
		processor.HasItem(l, c)(4032128) ||
		processor.HasItem(l, c)(4032129) {
		processor.SendPinkNotice(l, c)("ALREADY_HAVE_PROOF")
		return false
	}

	if processor.QuestStarted(l, c)(20611) ||
		processor.QuestStarted(l, c)(20612) ||
		processor.QuestStarted(l, c)(20613) ||
		processor.QuestStarted(l, c)(20614) ||
		processor.QuestStarted(l, c)(20615) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 913020300) == 0 {
			_map.KillAllMonsters(l)(c.WorldId(), c.ChannelId(), 913020300)
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(913020300, 0)
			_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 913020300, 9300294, 87, 88)
			return true
		}
		processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_ATTEMPTING_BOSS")
		return false
	} else {
		processor.SendPinkNotice(l, c)("CANNOT_ACCESS_HALL")
		return false
	}
}
