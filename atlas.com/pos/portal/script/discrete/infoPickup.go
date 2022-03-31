package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoPickup struct {
}

func (p InfoPickup) Name() string {
	return "infoPickup"
}

func (p InfoPickup) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(1035) {
		processor.ShowInfo(l, c)("UI/tutorial.img/21")
	}
	processor.BlockPortal(l, span, c)
	return true
}
