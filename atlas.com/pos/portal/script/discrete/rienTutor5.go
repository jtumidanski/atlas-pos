package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor5 struct {
}

func (p RienTutor5) Name() string {
	return "rienTutor5"
}

func (p RienTutor5) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.TalkGuide(l, c)("You're very close to town. I'll head over there first since I have some things to take care of. You take your time.")
	processor.BlockPortal(l, span, c)
	return false
}
