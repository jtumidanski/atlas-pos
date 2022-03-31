package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr3In1 struct {
}

func (p Jnr3In1) Name() string {
	return "jnr3_in1"
}

func (p Jnr3In1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out2"); err == nil && r.State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926110202, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
