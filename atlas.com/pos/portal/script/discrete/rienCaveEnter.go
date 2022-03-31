package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienCaveEnter struct {
}

func (p RienCaveEnter) Name() string {
	return "rienCaveEnter"
}

func (p RienCaveEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(21201) || processor.QuestStarted(l, c)(21302) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(140030000, 1)
		return true
	}
	processor.SendPinkNotice(l, c)("SOMETHING_BLOCKING_PORTAL")
	return false
}
