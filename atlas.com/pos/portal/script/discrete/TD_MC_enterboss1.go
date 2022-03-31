package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCEnterBoss1 struct {
}

func (p TDMCEnterBoss1) Name() string {
	return "TD_MC_enterboss1"
}

func (p TDMCEnterBoss1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	qp := processor.QuestProgressInt(l, c)(2330, 3300005) + processor.QuestProgressInt(l, c)(2330, 3300006) + processor.QuestProgressInt(l, c)(2330, 3300007)
	if processor.QuestStarted(l, c)(2330) && qp < 3 {
		processor.OpenNPC(l, c)(1300013)
	} else {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106021401, 1)
	}
	return true
}
