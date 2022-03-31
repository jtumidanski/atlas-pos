package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorMono1 struct {
}

func (p AranTutorMono1) Name() string {
	return "aranTutorMono1"
}

func (p AranTutorMono1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "mo2=o") {
		return false
	}
	processor.PlaySound(l, c)("Aran/balloon")
	processor.UpdateAreaInfo(l, c)(21002, "mo1=o;mo2=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon2")
	return true
}
