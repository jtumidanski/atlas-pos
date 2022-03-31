package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk01 struct {
}

func (p EvanTalk01) Name() string {
	return "evantalk01"
}

func (p EvanTalk01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "mo01=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon01")
	return true
}
