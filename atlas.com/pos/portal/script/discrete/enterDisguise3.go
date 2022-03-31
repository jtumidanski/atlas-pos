package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDisguise3 struct {
}

func (p EnterDisguise3) Name() string {
	return "enterDisguise3"
}

func (p EnterDisguise3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(20301) ||
		processor.QuestStarted(l, c)(20302) ||
		processor.QuestStarted(l, c)(20303) ||
		processor.QuestStarted(l, c)(20304) ||
		processor.QuestStarted(l, c)(20305) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010620) > 0 {
			processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if processor.HasItem(l, c)(4032103) {
			processor.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(108010620, "out00")
		return true
	}

	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(130010110, "out00")
	return true
}
