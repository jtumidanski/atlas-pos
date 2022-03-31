package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Kpq3 struct {
}

func (p Kpq3) Name() string {
	return "kpq3"
}

func (p Kpq3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(103000804)
	//if (eim.getProperty("4stageclear") != null) {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, target.getPortal("st00"))
	//	return true
	//}
	//else {
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
