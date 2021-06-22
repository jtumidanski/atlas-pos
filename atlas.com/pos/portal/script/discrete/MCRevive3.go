package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MCRevive3 struct {
}

func (p MCRevive3) Name() string {
	return "MCRevive3"
}

func (p MCRevive3) Enter(l logrus.FieldLogger, c script.Context) bool {
	portalId := uint32(0)
	//switch (pi.getTeam()) {
	//case 0:
	//	portal = 4
	//	break
	//case 1:
	//	portal = 3
	//	break
	//}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(980000301, portalId)
	return true
}
