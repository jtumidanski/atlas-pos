package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr3In0 struct {
}

func (p Jnr3In0) Name() string {
	return "jnr3_in0"
}

func (p Jnr3In0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out1").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926110201, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
