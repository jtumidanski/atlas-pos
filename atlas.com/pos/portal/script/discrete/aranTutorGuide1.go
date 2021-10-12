package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide1 struct {
}

func (p AranTutorGuide1) Name() string {
	return "aranTutorGuide1"
}

func (p AranTutorGuide1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)

	if script.ContainsAreaInfo(l, c)(21002, "chain=o") {
		return false
	}
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide2")
	script.SendPinkNotice(l, c)("ARAN_TUTORIAL_CONSECUTIVE")
	script.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;arr1=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
