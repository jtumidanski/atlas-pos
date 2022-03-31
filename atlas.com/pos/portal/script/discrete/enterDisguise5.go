package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDisguise5 struct {
}

func (p EnterDisguise5) Name() string {
	return "enterDisguise5"
}

func (p EnterDisguise5) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(20301) ||
		processor.QuestStarted(l, c)(20302) ||
		processor.QuestStarted(l, c)(20303) ||
		processor.QuestStarted(l, c)(20304) ||
		processor.QuestStarted(l, c)(20305) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010640) > 0 {
			processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if processor.HasItem(l, c)(4032105) {
			processor.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(108010640, "east00")
		return true
	}

	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(130020000, "east00")
	return true
}
