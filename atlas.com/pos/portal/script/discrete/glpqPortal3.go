package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal3 struct {
}

func (p GlpqPortal3) Name() string {
	return "glpqPortal3"
}

func (p GlpqPortal3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq3") < 5 || eim.getIntProperty("glpq3_p") < 5) {
	//		pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
	//		return false
	//	} else {
	//		pi.playPortalSound()
	//		pi.warp(610030400, 0)
	//		return true
	//	}
	//}
	//
	return false
}
