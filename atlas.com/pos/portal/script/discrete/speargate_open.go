package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type SpearGateOpen struct {
}

func (p SpearGateOpen) Name() string {
	return "speargate_open"
}

func (p SpearGateOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "speargate"); err == nil && r.State() == 4 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(990000401, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("NOT_OPEN_YET")
	return false
}
