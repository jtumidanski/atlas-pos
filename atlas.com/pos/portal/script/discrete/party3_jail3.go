package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Jail3 struct {
}

func (p Party3Jail3) Name() string {
	return "party3_jail3"
}

func (p Party3Jail3) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if(pi.getEventInstance().getIntProperty("statusStg8") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(920010930,0)
	//	return true
	//}
	//else {
	script.SendPinkNotice(l, c)("PIXIE_POWER_REMAINS")
	return false
	//}
}
