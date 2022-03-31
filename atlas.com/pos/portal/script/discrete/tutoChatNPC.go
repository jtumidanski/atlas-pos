package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TutoChatNPC struct {
}

func (p TutoChatNPC) Name() string {
	return "tutoChatNPC"
}

func (p TutoChatNPC) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasLevel30Character(l, span, c) {
		processor.OpenNPC(l, c)(2007)
	}
	processor.BlockPortal(l, span, c)
	return true
}
