package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
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
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "stonegate"); err == nil && r.State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000430, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("DOOR_IS_BLOCKED")
	return false
}
