package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanTalk50 struct {
}

func (p EvanTalk50) Name() string {
	return "evantalk50"
}

func (p EvanTalk50) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22014, "mo50=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22014, "mo30=o;mo40=o;mo41=o;mo50=o;mo42=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon50")
	return true
}
