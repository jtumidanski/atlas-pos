package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorArrow0 struct {
}

func (p AranTutorArrow0) Name() string {
	return "aranTutorArrow0"
}

func (p AranTutorArrow0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)

	if script.ContainsAreaInfo(l, c)(21002, "arr0=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "arr0=o;mo1=o;mo2=o;mo3=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow3")
	return true
}
