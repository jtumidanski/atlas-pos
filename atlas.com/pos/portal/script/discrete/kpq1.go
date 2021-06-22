package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Kpq1 struct {
}

func (p Kpq1) Name() string {
	return "kpq1"
}

func (p Kpq1) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(103000802)
	//if (eim.getProperty("2stageclear") != null) {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, target.getPortal("st00"))
	//	return true
	//}
	//else {
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
