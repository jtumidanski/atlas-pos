package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoSkill struct {
}

func (p InfoSkill) Name() string {
	return "infoSkill"
}

func (p InfoSkill) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestCompleted(l, c)(1035) {
		script.ShowInfo(l, c)("UI/tutorial.img/23")
	}
	script.BlockPortal(l, span, c)
	return true
}
