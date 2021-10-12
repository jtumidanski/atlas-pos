package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type StoneGateOpen struct {
}

func (p StoneGateOpen) Name() string {
	return "stonegate_open"
}

func (p StoneGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "stonegate").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(990000430, 0)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_IS_BLOCKED")
	return false
}
