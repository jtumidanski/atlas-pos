package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr3Out struct {
}

func (p Jnr3Out) Name() string {
	return "jnr3_out"
}

func (p Jnr3Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "jnr3_out3"); err == nil && r.State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926110203, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
