package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TutoChatNPC struct {
}

func (p TutoChatNPC) Name() string {
	return "tutoChatNPC"
}

func (p TutoChatNPC) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasLevel30Character(l, span, c) {
		script.OpenNPC(l, c)(2007)
	}
	script.BlockPortal(l, span, c)
	return true
}
