package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj6Out struct {
}

func (p Rnj6Out) Name() string {
	return "rnj6_out"
}

func (p Rnj6Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj6_out").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926100300, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
