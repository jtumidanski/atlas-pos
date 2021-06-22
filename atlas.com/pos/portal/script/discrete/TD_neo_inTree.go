package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TDNeoInTree struct {
}

func (p TDNeoInTree) Name() string {
	return "TD_neo_inTree"
}

func (p TDNeoInTree) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventManager nex = pi.getEventManager("GuardianNex")
	//if (nex == null) {
	//	pi.sendPinkNotice("GUARDIAN_NEX_ERROR")
	//	return false
	//}
	//
	//int[] quests = [3719, 3724, 3730, 3736, 3742, 3748]
	//int[] mobs = [7120100, 7120101, 7120102, 8120100, 8120101, 8140510]
	//
	//for (int i = 0; i < quests.length; i++) {
	//if (pi.isQuestActive(quests[i])) {
	//if (pi.getQuestProgressInt(quests[i], mobs[i]) != 0) {
	//pi.sendPinkNotice("GUARDIAN_NEX_ALREADY_FACED")
	//return false
	//}
	//
	//if (!nex.startInstance(i, pi.getPlayer())) {
	//pi.sendPinkNotice("GUARDIAN_NEX_SOMEONE_ALREADY_FACING")
	//return false
	//} else {
	//pi.playPortalSound()
	//return true
	//}
	//}
	//}

	script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
