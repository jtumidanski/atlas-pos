package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OBoat2 struct {
}

func (p OBoat2) Name() string {
	return "OBoat2"
}

func (p OBoat2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(200090010, 5)
	//if (pi.getPlayer().getClient().getChannelServer().getEventSM().getEventManager("Boats").getProperty("haveBalrog") == "true") {
	//	pi.changeMusic("Bgm04/ArabPirate")
	//}
	return true
}
