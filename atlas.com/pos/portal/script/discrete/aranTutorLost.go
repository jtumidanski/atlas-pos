package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorLost struct {
}

func (p AranTutorLost) Name() string {
	return "aranTutorLost"
}

func (p AranTutorLost) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)

	if processor.ContainsAreaInfo(l, c)(21002, "fin=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;fin=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	processor.ShowIntro(l, c)("Effect/Direction1.img/aranTutorial/ClickChild")
	return true
}
