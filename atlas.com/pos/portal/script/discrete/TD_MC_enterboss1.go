package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCEnterBoss1 struct {
}

func (p TDMCEnterBoss1) Name() string {
	return "TD_MC_enterboss1"
}

func (p TDMCEnterBoss1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	qp := script.QuestProgressInt(l, c)(2330, 3300005) + script.QuestProgressInt(l, c)(2330, 3300006) + script.QuestProgressInt(l, c)(2330, 3300007)
	if script.QuestStarted(l, c)(2330) && qp < 3 {
		script.OpenNPC(l, c)(1300013)
	} else {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(106021401, 1)
	}
	return true
}
