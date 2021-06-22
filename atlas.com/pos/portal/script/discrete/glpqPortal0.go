package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal0 struct {
}

func (p GlpqPortal0) Name() string {
	return "glpqPortal0"
}

func (p GlpqPortal0) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("glpq1") == 0) {
	//	pi.sendPinkNotice("PATH_BLOCKED")
	//	return false
	//
	//} else {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(610030200, 0)
	return true
	//}
}
