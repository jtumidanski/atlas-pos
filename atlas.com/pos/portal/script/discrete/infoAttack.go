package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoAttack struct {
}

func (p InfoAttack) Name() string {
	return "infoAttack"
}

func (p InfoAttack) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(1035) {
		script.ShowInfo(l, c)("UI/tutorial.img/20")
	}
	script.BlockPortal(l, span, c)
	return true
}
