package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDollWay struct {
}

func (p EnterDollWay) Name() string {
	return "enterDollWay"
}

func (p EnterDollWay) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(20730) || processor.QuestCompleted(l, c)(21734) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(105070300, 3)
		return true
	}
	if processor.QuestStarted(l, c)(21734) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(910510100, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("OMINOUS_POWER")
	return false
}
