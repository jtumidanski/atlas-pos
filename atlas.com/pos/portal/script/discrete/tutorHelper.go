package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TutorHelper struct {
}

func (p TutorHelper) Name() string {
	return "tutorHelper"
}

func (p TutorHelper) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SpawnGuide(l, c)
	processor.TalkGuide(l, c)("Welcome to Maple World! I'm Mimo. I'm in charge of guiding you until you reach Lv. 10 and become a Knight-In-Training. Double-click for further information!")
	processor.BlockPortal(l, span, c)
	return true
}
