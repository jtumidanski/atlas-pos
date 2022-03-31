package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDollcave struct {
}

func (p EnterDollcave) Name() string {
	return "enterDollcave"
}

func (p EnterDollcave) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(20730) || processor.QuestCompleted(l, c)(21734) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(105040201, 2)
		return true
	}
	processor.OpenNPCWithScript(l, c)(1063011, "PupeteerPassword")
	return false
}
