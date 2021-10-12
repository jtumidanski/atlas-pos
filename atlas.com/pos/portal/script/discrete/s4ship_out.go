package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4ShipOut struct {
}

func (p S4ShipOut) Name() string {
	return "s4ship_out"
}

func (p S4ShipOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//int exit = pi.getEventInstance().getIntProperty("canLeave")
	//if (exit == 0) {
	//	pi.sendPinkNotice("WAIT_TO_LEAVE")
	//	return false
	//} else if (exit == 2) {
	//	pi.playPortalSound()
	//	pi.warp(912010200)
	//	return true
	//} else {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, span, c)(120000101)
	return true
	//}
}
