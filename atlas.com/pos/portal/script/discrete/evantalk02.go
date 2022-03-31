package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk02 struct {
}

func (p EvanTalk02) Name() string {
	return "evantalk02"
}

func (p EvanTalk02) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "mo02=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o;mo02=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon02")
	return true
}
