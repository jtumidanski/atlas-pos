package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr6Out struct {
}

func (p Jnr6Out) Name() string {
	return "jnr6_out"
}

func (p Jnr6Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr6_out").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926110300, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
