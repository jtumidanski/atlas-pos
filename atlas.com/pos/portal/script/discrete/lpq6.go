package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Lpq6 struct {
}

func (p Lpq6) Name() string {
	return "lpq6"
}

func (p Lpq6) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(922010800)
	//if (eim.getProperty("7stageclear") != null) {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, target.getPortal("st00"))
	//	return true
	//} else {
		return false
	//}
}
