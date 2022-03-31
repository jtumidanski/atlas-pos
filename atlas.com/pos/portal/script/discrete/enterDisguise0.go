package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDisguise0 struct {
}

func (p EnterDisguise0) Name() string {
	return "enterDisguise0"
}

func (p EnterDisguise0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(20301) ||
		processor.QuestStarted(l, c)(20302) ||
		processor.QuestStarted(l, c)(20303) ||
		processor.QuestStarted(l, c)(20304) ||
		processor.QuestStarted(l, c)(20305) {
		if processor.HasItem(l, c)(4032179) {
			processor.PlayPortalSound(l, c)
			processor.WarpByName(l, span, c)(130010000, "east00")
			return true
		}
		processor.SendPinkNotice(l, c)("NEED_PERMIT")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(130010000, "east00")
	return true
}
