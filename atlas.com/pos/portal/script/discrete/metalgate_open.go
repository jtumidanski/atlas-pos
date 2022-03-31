package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MetalGateOpen struct {
}

func (p MetalGateOpen) Name() string {
	return "metalgate_open"
}

func (p MetalGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "metalgate"); err == nil && r.State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000431, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
