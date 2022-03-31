package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk10 struct {
}

func (p EvanTalk10) Name() string {
	return "evantalk10"
}

func (p EvanTalk10) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "mo10=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o;mo10=0;mo02=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon10")
	return true
}
