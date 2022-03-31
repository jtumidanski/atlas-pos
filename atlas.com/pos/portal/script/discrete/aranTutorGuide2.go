package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide2 struct {
}

func (p AranTutorGuide2) Name() string {
	return "aranTutorGuide2"
}

func (p AranTutorGuide2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "cmd=o") {
		return false
	}
	processor.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide3")
	processor.SendPinkNotice(l, c)("ARAN_TUTORIAL_COMMAND")
	processor.UpdateAreaInfo(l, c)(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
