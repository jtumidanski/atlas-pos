package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type KingGate2Open struct {
}

func (p KingGate2Open) Name() string {
	return "kinggate2_open"
}

func (p KingGate2Open) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "kinggate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(990000900, 2)
		//if (pi.getPlayer().getEventInstance().getProperty("boss") != null && pi.getPlayer().getEventInstance().getProperty("boss") == "true") {
		//	pi.changeMusic("Bgm10/Eregos")
		//}
		return true
	}
	script.SendPinkNotice(l, c)("KING_GATE")
	return false
}
