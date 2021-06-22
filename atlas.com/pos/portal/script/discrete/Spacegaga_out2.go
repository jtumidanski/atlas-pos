package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type SpaceGagaOut2 struct {
}

func (p SpaceGagaOut2) Name() string {
	return "Spacegaga_out2"
}

func (p SpaceGagaOut2) Enter(l logrus.FieldLogger, c script.Context) bool {
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
