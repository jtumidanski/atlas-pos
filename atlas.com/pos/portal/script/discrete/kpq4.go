package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Kpq4 struct {
}

func (p Kpq4) Name() string {
	return "kpq4"
}

func (p Kpq4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(103000805)
	//if (eim.getProperty("5stageclear") != null) {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, target.getPortal("st00"))
	//	return true
	//}
	//else {
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
