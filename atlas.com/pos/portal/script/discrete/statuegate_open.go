package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type StatueGateOpen struct {
}

func (p StatueGateOpen) Name() string {
	return "statuegate_open"
}

func (p StatueGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "statuegate").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000301, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("GATE_IS_CLOSED")
	return false
}
