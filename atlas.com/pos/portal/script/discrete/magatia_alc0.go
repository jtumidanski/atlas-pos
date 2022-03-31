package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MagatiaAlc0 struct {
}

func (p MagatiaAlc0) Name() string {
	return "magatia_alc0"
}

func (p MagatiaAlc0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(3309) || processor.HasItem(l, c)(4031708) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(261020700, "down00")
	} else {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(926120000, "out00")
	}
	return true
}
