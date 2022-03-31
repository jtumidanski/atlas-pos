package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanRoom0 struct {
}

func (p EvanRoom0) Name() string {
	return "evanRoom0"
}

func (p EvanRoom0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22014, "mo30=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22014, "mo30=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon30")
	return true
}
