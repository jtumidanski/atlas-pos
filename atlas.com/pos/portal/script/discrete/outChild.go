package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutChild struct {
}

func (p OutChild) Name() string {
	return "outChild"
}

func (p OutChild) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(21001) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(914000220, 2)
		return true
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(914000400, 2)
	return true
}
