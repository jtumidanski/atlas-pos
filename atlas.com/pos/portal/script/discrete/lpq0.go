package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Lpq0 struct {
}

func (p Lpq0) Name() string {
	return "lpq0"
}

func (p Lpq0) Enter(l logrus.FieldLogger, c script.Context) bool {
	//int nextMap = 922010200
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//String avail = eim.getProperty("1stageclear")
	//if (avail == null) {
	script.SendPinkNotice(l, c)("SEAL_BLOCKING_DOOR")
	return false
	//}
	//else {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
}
