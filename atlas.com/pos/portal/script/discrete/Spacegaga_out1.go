package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type SpaceGagaOut1 struct {
}

func (p SpaceGagaOut1) Name() string {
	return "Spacegaga_out1"
}

func (p SpaceGagaOut1) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//int fc = eim.getIntProperty("falls")
	//
	//if (fc >= 3) {
	//	pi.playPortalSound()
	//	pi.warp(922240200, 0)
	//} else {
	//	eim.setIntProperty("falls", fc + 1)
	//	pi.playPortalSound()
	//	pi.warp(pi.getMapId(), 0)
	//}

	return true
}
