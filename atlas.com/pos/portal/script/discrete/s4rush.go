package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Rush struct {
}

func (p S4Rush) Name() string {
	return "s4rush"
}

func (p S4Rush) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(6110) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(910500100, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
