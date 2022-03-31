package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3Jail3 struct {
}

func (p Party3Jail3) Name() string {
	return "party3_jail3"
}

func (p Party3Jail3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if(pi.getEventInstance().getIntProperty("statusStg8") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(920010930,0)
	//	return true
	//}
	//else {
	processor.SendPinkNotice(l, c)("PIXIE_POWER_REMAINS")
	return false
	//}
}
