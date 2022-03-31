package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoSkill struct {
}

func (p InfoSkill) Name() string {
	return "infoSkill"
}

func (p InfoSkill) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(1035) {
		processor.ShowInfo(l, c)("UI/tutorial.img/23")
	}
	processor.BlockPortal(l, span, c)
	return true
}
