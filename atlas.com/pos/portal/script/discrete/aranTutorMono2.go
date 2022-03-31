package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorMono2 struct {
}

func (p AranTutorMono2) Name() string {
	return "aranTutorMono2"
}

func (p AranTutorMono2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "mo3=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(21002, "mo1=o;mo2=o;mo3=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon3")
	return true
}
