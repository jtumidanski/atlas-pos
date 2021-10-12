package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterMCave struct {
}

func (p EnterMCave) Name() string {
	return "enterMCave"
}

func (p EnterMCave) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(21201) {
		for i := uint32(108000700); i < 108000709; i++ {
			if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), i) > 0 && _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), i+10) > 0 {
				continue
			}
			script.PlayPortalSound(l, c)
			script.WarpByName(l, span, c)(i, "out00")
			script.SetQuestProgress(l, c)(21202, 21203, 0)
			return true
		}
		script.SendPinkNotice(l, c)("MIRROR_IS_BLANK")
		return false
	}

	if script.QuestStarted(l, c)(21302) && !script.QuestCompleted(l, c)(21303) {
		if len(_map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), 108010701)) > 0 || len(_map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), 108010702)) > 0 {
			script.SendPinkNotice(l, c)("MIRROR_IS_BLANK")
			return false
		}
		_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 108010702, 9001013, -210, 454)
		script.PlayPortalSound(l, c)
		script.SetQuestProgress(l, c)(21303, 21203, 1)
		script.WarpByName(l, span, c)(108010701, "out00")
		return true
	}

	script.SendPinkNotice(l, c)("MIRROR_ALREADY_PASSED")
	return false
}
