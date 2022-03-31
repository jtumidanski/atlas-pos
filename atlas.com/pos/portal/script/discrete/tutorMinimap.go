package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TutorMiniMap struct {
}

func (p TutorMiniMap) Name() string {
	return "tutorMinimap"
}

func (p TutorMiniMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.GuideHint(l, c)(1)
	processor.BlockPortal(l, span, c)
	return true
}
