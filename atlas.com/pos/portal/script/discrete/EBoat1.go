package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EBoat1 struct {
}

func (p EBoat1) Name() string {
	return "EBoat1"
}

func (p EBoat1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(200090000, 4)
	//if (pi.getPlayer().getClient().getChannelServer().getEventSM().getEventManager("Boats").getProperty("haveBalrog") == "true") {
	//	pi.changeMusic("Bgm04/ArabPirate")
	//}
	return true
}
