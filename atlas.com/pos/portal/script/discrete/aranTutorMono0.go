package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorMono0 struct {
}

func (p AranTutorMono0) Name() string {
	return "aranTutorMono0"
}

func (p AranTutorMono0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "mo1=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(21002, "mo1=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon1")
	return true
}
