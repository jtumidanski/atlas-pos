package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TristanEnter struct {
}

func (p TristanEnter) Name() string {
	return "tristanEnter"
}

func (p TristanEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestCompleted(l, c)(2238) {
		processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(105100101, "in00")
	return true
}
