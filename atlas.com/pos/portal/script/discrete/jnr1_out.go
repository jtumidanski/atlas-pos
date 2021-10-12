package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr1Out struct {
}

func (p Jnr1Out) Name() string {
	return "jnr1_out"
}

func (p Jnr1Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("statusStg2") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(926110100, 0) //next
	//	return true
	//} else {
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
