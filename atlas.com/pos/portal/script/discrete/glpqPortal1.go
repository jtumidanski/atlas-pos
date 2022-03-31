package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal1 struct {
}

func (p GlpqPortal1) Name() string {
	return "glpqPortal1"
}

func (p GlpqPortal1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq2") == 5) {
	//script.PlayPortalSound(l, c)
	//script.WarpById(l, span, c)(610030300, 0)
	//return true
	//	} else {
	processor.SendPinkNotice(l, c)("PORTAL_NOT_ACTIVE")
	return false
	//	}
}
