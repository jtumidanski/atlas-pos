package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HorntaleBR struct {
}

func (p HorntaleBR) Name() string {
	return "hontale_BR"
}

func (p HorntaleBR) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() == 240060000 {
		//if (pi.getEventInstance().getIntProperty("defeatedHead") >= 1) {
		//	pi.playPortalSound()
		//	pi.warp(240060100, 0)
		//	return true
		//} else {
		processor.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//}
	} else if c.MapId() == 240060100 {
		//if (pi.getEventInstance().getIntProperty("defeatedHead") >= 2) {
		//	pi.playPortalSound()
		//	pi.warp(240060200, 0)
		//	return true
		//} else {
		processor.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//}
	}
	return false
}
