package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutPepeKing struct {
}

func (p OutPepeKing) Name() string {
	return "out_pepeking"
}

func (p OutPepeKing) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	eim.stopEventTimer()
	//	eim.dispose()
	//}

	qp := processor.QuestProgressInt(l, c)(2330, 3300005) + processor.QuestProgressInt(l, c)(2330, 3300006) + processor.QuestProgressInt(l, c)(2330, 3300007)
	if qp != 3 || processor.HasItem(l, c)(4032388) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106021400, 2)
		return true
	}

	if !processor.CanHold(l, c)(4032388, 1) {
		processor.SendPinkNotice(l, c)("INVENTORY_FULL")
		return false
	}

	processor.SendPinkNotice(l, c)("PEPE_KING_DROP")
	processor.GainItem(l, c)(4032388, 1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(106021400, 2)
	return true
}
