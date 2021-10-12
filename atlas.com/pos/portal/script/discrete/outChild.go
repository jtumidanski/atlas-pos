package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutChild struct {
}

func (p OutChild) Name() string {
	return "outChild"
}

func (p OutChild) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestStarted(l, c)(21001) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(914000220, 2)
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(914000400, 2)
	return true
}
