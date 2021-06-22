package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutPepeKing struct {
}

func (p OutPepeKing) Name() string {
	return "out_pepeking"
}

func (p OutPepeKing) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	eim.stopEventTimer()
	//	eim.dispose()
	//}

	qp := script.QuestProgressInt(l, c)(2330, 3300005) + script.QuestProgressInt(l, c)(2330, 3300006) + script.QuestProgressInt(l, c)(2330, 3300007)
	if qp != 3 || script.HasItem(l, c)(4032388) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(106021400, 2)
		return true
	}

	if !script.CanHold(l, c)(4032388, 1) {
		script.SendPinkNotice(l, c)("INVENTORY_FULL")
		return false
	}

	script.SendPinkNotice(l, c)("PEPE_KING_DROP")
	script.GainItem(l, c)(4032388, 1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(106021400, 2)
	return true
}
