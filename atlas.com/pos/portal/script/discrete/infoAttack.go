package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoAttack struct {
}

func (p InfoAttack) Name() string {
	return "infoAttack"
}

func (p InfoAttack) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(1035) {
		processor.ShowInfo(l, c)("UI/tutorial.img/20")
	}
	processor.BlockPortal(l, span, c)
	return true
}
