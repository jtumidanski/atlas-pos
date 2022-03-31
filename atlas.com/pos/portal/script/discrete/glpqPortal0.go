package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal0 struct {
}

func (p GlpqPortal0) Name() string {
	return "glpqPortal0"
}

func (p GlpqPortal0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("glpq1") == 0) {
	//	pi.sendPinkNotice("PATH_BLOCKED")
	//	return false
	//
	//} else {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(610030200, 0)
	return true
	//}
}
