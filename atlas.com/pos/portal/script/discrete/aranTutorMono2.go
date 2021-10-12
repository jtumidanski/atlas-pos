package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorMono2 struct {
}

func (p AranTutorMono2) Name() string {
	return "aranTutorMono2"
}

func (p AranTutorMono2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)

	if script.ContainsAreaInfo(l, c)(21002, "mo3=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "mo1=o;mo2=o;mo3=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon3")
	return true
}
