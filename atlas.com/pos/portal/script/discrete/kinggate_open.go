package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type KingGateOpen struct {
}

func (p KingGateOpen) Name() string {
	return "kinggate_open"
}

func (p KingGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "kinggate"); err == nil && r.State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000900, 1)
		//if (pi.getPlayer().getEventInstance().getProperty("boss") != null && pi.getPlayer().getEventInstance().getProperty("boss") == "true") {
		//	pi.changeMusic("Bgm10/Eregos")
		//}
		return true
	}
	processor.SendPinkNotice(l, c)("DOOR_IS_CLOSED")
	return false
}
