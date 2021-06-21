package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterRider struct {
}

func (p EnterRider) Name() string {
	return "enterRider"
}

func (p EnterRider) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(21610) &&
		script.HasItem(l, c)(4001193) {
		script.SendPinkNotice(l, c)("2ND_MOUNT_QUEST_CLOSED")
		return false
		//EventManager em = pi.getEventManager("Aran_2ndmount")
		//if (em == null) {
		//	pi.sendPinkNotice("2ND_MOUNT_QUEST_CLOSED")
		//	return false
		//} else {
		//	if (!em.startInstance(pi.getPlayer())) {
		//		pi.sendPinkNotice("2ND_MOUNT_QUEST_SOMEONE_ALREADY_IN")
		//		return false
		//	} else {
		//		pi.playPortalSound()
		//		return true
		//	}
		//}
	}

	script.SendPinkNotice(l, c)("2ND_MOUNT_QUEST_REQUIREMENT")
	return false
}
