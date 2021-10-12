package discrete

import (
	"atlas-pos/portal/script"
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
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "kinggate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(990000900, 1)
		//if (pi.getPlayer().getEventInstance().getProperty("boss") != null && pi.getPlayer().getEventInstance().getProperty("boss") == "true") {
		//	pi.changeMusic("Bgm10/Eregos")
		//}
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_IS_CLOSED")
	return false
}
