package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide2 struct {
}

func (p AranTutorGuide2) Name() string {
	return "aranTutorGuide2"
}

func (p AranTutorGuide2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)

	if script.ContainsAreaInfo(l, c)(21002, "cmd=o") {
		return false
	}
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide3")
	script.SendPinkNotice(l, c)("ARAN_TUTORIAL_COMMAND")
	script.UpdateAreaInfo(l, c)(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
