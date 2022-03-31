package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutDarkEreb struct {
}

func (p OutDarkEreb) Name() string {
	return "outDarkEreb"
}

func (p OutDarkEreb) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestCompleted(l, c)(20407) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(924010100, 0)
		return true
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(924010200, 0)
	return true
}
