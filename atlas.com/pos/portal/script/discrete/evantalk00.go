package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk00 struct {
}

func (p EvanTalk00) Name() string {
	return "evantalk00"
}

func (p EvanTalk00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "mo00=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22013, "mo00=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon00")
	return true
}
