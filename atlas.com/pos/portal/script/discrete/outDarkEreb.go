package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutDarkEreb struct {
}

func (p OutDarkEreb) Name() string {
	return "outDarkEreb"
}

func (p OutDarkEreb) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestCompleted(l, c)(20407) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(924010100, 0)
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(924010200, 0)
	return true
}
