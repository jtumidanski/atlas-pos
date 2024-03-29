package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorArrow1 struct {
}

func (p AranTutorArrow1) Name() string {
	return "aranTutorArrow1"
}

func (p AranTutorArrow1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "arr1=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;arr1=o;mo1=o;mo2=o;mo3=o;mo4=o")
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
