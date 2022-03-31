package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterRider struct {
}

func (p EnterRider) Name() string {
	return "enterRider"
}

func (p EnterRider) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(21610) &&
		processor.HasItem(l, c)(4001193) {
		processor.SendPinkNotice(l, c)("2ND_MOUNT_QUEST_CLOSED")
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

	processor.SendPinkNotice(l, c)("2ND_MOUNT_QUEST_REQUIREMENT")
	return false
}
