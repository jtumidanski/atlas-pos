package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GuyFawkes0Esc struct {
}

func (p GuyFawkes0Esc) Name() string {
	return "guyfawkes0_esc"
}

func (p GuyFawkes0Esc) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("statusStg1") == 1) {
	//	pi.playPortalSound()
	//	pi.warp(674030200, 0)
	//	return true
	//} else {
	script.SendPinkNotice(l, c)("TUNNEL_BLOCKED")
	return false
	//}
}
