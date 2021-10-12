package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoReactor struct {
}

func (p InfoReactor) Name() string {
	return "infoReactor"
}

func (p InfoReactor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestCompleted(l, c)(1008) {
		script.ShowInfo(l, c)("UI/tutorial.img/22")
	} else if script.QuestCompleted(l, c)(1020) {
		script.ShowInfo(l, c)("UI/tutorial.img/27")
	}
	script.BlockPortal(l, span, c)
	return true
}
