package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal4 struct {
}

func (p GlpqPortal4) Name() string {
	return "glpqPortal4"
}

func (p GlpqPortal4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq4") < 5) {
	//		pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
	//		return false
	//	} else {
	//		pi.playPortalSound()
	//		pi.warp(610030500, 0)
	//		return true
	//	}
	//}
	//
	return false
}
