package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj2Out struct {
}

func (p Rnj2Out) Name() string {
	return "rnj2_out"
}

func (p Rnj2Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "statusStg3").State() == 3 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(926100200, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
