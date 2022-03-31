package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GhostGateOpen struct {
}

func (p GhostGateOpen) Name() string {
	return "ghostgate_open"
}

func (p GhostGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "ghostgate").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000800, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
