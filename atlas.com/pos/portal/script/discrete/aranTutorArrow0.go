package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorArrow0 struct {
}

func (p AranTutorArrow0) Name() string {
	return "aranTutorArrow0"
}

func (p AranTutorArrow0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "arr0=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(21002, "arr0=o;mo1=o;mo2=o;mo3=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow3")
	return true
}
