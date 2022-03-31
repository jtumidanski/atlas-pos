package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide0 struct {
}

func (p AranTutorGuide0) Name() string {
	return "aranTutorGuide0"
}

func (p AranTutorGuide0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "normal=o") {
		return false
	}
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide1")
	processor.SendPinkNotice(l, c)("ARAN_TUTORIAL_REGULAR_ATTACK")
	processor.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o")
	return true
}
