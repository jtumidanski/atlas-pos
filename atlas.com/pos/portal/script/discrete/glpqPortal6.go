package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal6 struct {
}

func (p GlpqPortal6) Name() string {
	return "glpqPortal6"
}

func (p GlpqPortal6) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq6") < 3) {
	//		pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
	//		return false
	//	} else {
	//		pi.playPortalSound()
	//		pi.warp(610030700, 0)
	//		return true
	//	}
	//}

	return false
}
