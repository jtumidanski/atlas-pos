package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3Room8 struct {
}

func (p Party3Room8) Name() string {
	return "party3_room8"
}

func (p Party3Room8) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("statusStg8") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(920011000, 0)
	//	return true
	//} else {
	processor.SendPinkNotice(l, c)("PIXIE_POWER_REMAINS")
	return false
	//}
}
