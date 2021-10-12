package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr1Pt00 struct {
}

func (p Jnr1Pt00) Name() string {
	return "jnr1_pt00"
}

func (p Jnr1Pt00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("statusStg1") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(926110001, 0) //next
	//	return true
	//} else {
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
