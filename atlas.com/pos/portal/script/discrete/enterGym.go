package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterGym struct {
}

func (p EnterGym) Name() string {
	return "enterGym"
}

func (p EnterGym) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(21701) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(914010000, 1)
		return true
	}
	if processor.QuestStarted(l, c)(21702) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(914010100, 1)
		return true
	}
	if processor.QuestStarted(l, c)(21703) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(914010200, 1)
		return true
	}
	processor.SendPinkNotice(l, c)("PUO_LESSON_REQUIREMENT")
	return false
}
