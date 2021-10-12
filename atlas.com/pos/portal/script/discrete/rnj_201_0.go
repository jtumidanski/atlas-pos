package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj2010 struct {
}

func (p Rnj2010) Name() string {
	return "rnj_201_0"
}

func (p Rnj2010) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj31_out").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(926100200, 1)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_NOT_YET_OPENED")
	return false
}
