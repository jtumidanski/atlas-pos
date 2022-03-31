package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type WaterGateOpen struct {
}

func (p WaterGateOpen) Name() string {
	return "watergate_open"
}

func (p WaterGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "watergate").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000600, 1)
		return true
	}
	processor.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
