package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3Jail1 struct {
}

func (p Party3Jail1) Name() string {
	return "party3_jail1"
}

func (p Party3Jail1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if(pi.getEventInstance().getIntProperty("statusStg8") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(920010910,0)
	//	return true
	//}
	//else {
	script.SendPinkNotice(l, c)("PIXIE_POWER_REMAINS")
	return false
	//}
}
