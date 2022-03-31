package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoReactor struct {
}

func (p InfoReactor) Name() string {
	return "infoReactor"
}

func (p InfoReactor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(1008) {
		processor.ShowInfo(l, c)("UI/tutorial.img/22")
	} else if processor.QuestCompleted(l, c)(1020) {
		processor.ShowInfo(l, c)("UI/tutorial.img/27")
	}
	processor.BlockPortal(l, span, c)
	return true
}
