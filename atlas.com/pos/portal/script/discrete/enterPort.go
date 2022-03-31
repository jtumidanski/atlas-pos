package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterPort struct {
}

func (p EnterPort) Name() string {
	return "enterPort"
}

func (p EnterPort) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(21301) &&
		processor.QuestProgressInt(l, c)(21301, 9001013) == 0 {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010700) != 0 {
			processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_CHALLENGING_THIEF_CROW")
			return false
		}
		_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 108010700, 9001013, 2732, 3)
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(108010700, "west00")
		return true
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(140020300, 1)
	return true
}
