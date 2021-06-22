package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Lpq5 struct {
}

func (p Lpq5) Name() string {
	return "lpq5"
}

func (p Lpq5) Enter(l logrus.FieldLogger, c script.Context) bool {
	//int nextMap = 922010700
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//// only let people through if the eim is ready
	//String avail = eim.getProperty("5stageclear")
	//if (avail == null) {
	//	// can't go thru eh?
	script.SendPinkNotice(l, c)("SEAL_BLOCKING_DOOR")
	return false
	//} else {
	//	if (eim.getProperty("6stageclear") == null) {
	//		eim.setProperty("6stageclear", "true")
	//	}
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
}
