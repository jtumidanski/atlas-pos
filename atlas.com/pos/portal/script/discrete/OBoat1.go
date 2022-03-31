package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OBoat1 struct {
}

func (p OBoat1) Name() string {
	return "OBoat1"
}

func (p OBoat1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(200090010, 4)
	//if (pi.getPlayer().getClient().getChannelServer().getEventSM().getEventManager("Boats").getProperty("haveBalrog") == "true") {
	//	pi.changeMusic("Bgm04/ArabPirate")
	//}
	return true
}
