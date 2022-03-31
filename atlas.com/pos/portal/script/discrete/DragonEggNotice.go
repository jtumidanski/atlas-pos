package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DragonEggNotice struct {
}

func (p DragonEggNotice) Name() string {
	return "DrnEggNotice"
}

func (p DragonEggNotice) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22014, "egg=o") {
		return false
	}
	processor.UpdateAreaInfo(l, c)(22014, "egg=o;mo30=o;mo40=o;mo41=o;mo50=o;mo42=o;mo60=o")
	processor.ShowInfo(l, c)("UI/tutorial/evan/8/0")
	return true
}
