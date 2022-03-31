package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoMiniMap struct {
}

func (p InfoMiniMap) Name() string {
	return "infoMinimap"
}

func (p InfoMiniMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(1031) {
		processor.ShowInfo(l, c)("UI/tutorial.img/25")
	}
	processor.BlockPortal(l, span, c)
	return true
}
