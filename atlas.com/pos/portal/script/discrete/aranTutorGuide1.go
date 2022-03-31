package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide1 struct {
}

func (p AranTutorGuide1) Name() string {
	return "aranTutorGuide1"
}

func (p AranTutorGuide1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "chain=o") {
		return false
	}
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide2")
	processor.SendPinkNotice(l, c)("ARAN_TUTORIAL_CONSECUTIVE")
	processor.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;arr1=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
