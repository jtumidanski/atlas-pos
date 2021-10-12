package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorLost struct {
}

func (p AranTutorLost) Name() string {
	return "aranTutorLost"
}

func (p AranTutorLost) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)

	if script.ContainsAreaInfo(l, c)(21002, "fin=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;fin=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	script.ShowIntro(l, c)("Effect/Direction1.img/aranTutorial/ClickChild")
	return true
}
