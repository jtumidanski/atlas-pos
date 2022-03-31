package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk20 struct {
}

func (p EvanTalk20) Name() string {
	return "evantalk20"
}

func (p EvanTalk20) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "mo20=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o;mo11=o;mo20=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon20")
	return true
}
