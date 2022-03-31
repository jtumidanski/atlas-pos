package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanGarden1 struct {
}

func (p EvanGarden1) Name() string {
	return "evanGarden1"
}

func (p EvanGarden1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(22008) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(100030103, "west00")
	} else {
		processor.SendPinkNotice(l, c)("CANNOT_ENTER_BACKYARD_WITHOUT")
	}
	return true
}
