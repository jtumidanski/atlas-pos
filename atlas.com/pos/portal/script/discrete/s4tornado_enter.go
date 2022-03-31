package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4TornadoEnter struct {
}

func (p S4TornadoEnter) Name() string {
	return "s4tornado_enter"
}

func (p S4TornadoEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(6230) || processor.QuestStarted(l, c)(6231) || processor.HasItem(l, c)(4001110) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 922020200) == 0 {
			_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 922020200)
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(922020200, 0)
			return true
		}
		processor.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}
	return false
}
