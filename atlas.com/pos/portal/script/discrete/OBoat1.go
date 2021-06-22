package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OBoat1 struct {
}

func (p OBoat1) Name() string {
	return "OBoat1"
}

func (p OBoat1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(200090010, 4)
	//if (pi.getPlayer().getClient().getChannelServer().getEventSM().getEventManager("Boats").getProperty("haveBalrog") == "true") {
	//	pi.changeMusic("Bgm04/ArabPirate")
	//}
	return true
}
