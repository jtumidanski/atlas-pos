package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal1 struct {
}

func (p GlpqPortal1) Name() string {
	return "glpqPortal1"
}

func (p GlpqPortal1) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq2") == 5) {
	//script.PlayPortalSound(l, c)
	//script.WarpById(l, c)(610030300, 0)
	//return true
	//	} else {
	script.SendPinkNotice(l, c)("PORTAL_NOT_ACTIVE")
	return false
	//	}
}
