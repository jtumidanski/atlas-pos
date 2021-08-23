package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OBoat2 struct {
}

func (p OBoat2) Name() string {
	return "OBoat2"
}

func (p OBoat2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(200090010, 5)
	//if (pi.getPlayer().getClient().getChannelServer().getEventSM().getEventManager("Boats").getProperty("haveBalrog") == "true") {
	//	pi.changeMusic("Bgm04/ArabPirate")
	//}
	return true
}