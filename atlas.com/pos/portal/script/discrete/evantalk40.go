package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk40 struct {
}

func (p EvanTalk40) Name() string {
	return "evantalk40"
}

func (p EvanTalk40) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22014, "mo40=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22014, "mo30=o;mo40=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon40")
	return true
}
