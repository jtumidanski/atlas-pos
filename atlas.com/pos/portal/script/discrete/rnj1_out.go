package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj1Out struct {
}

func (p Rnj1Out) Name() string {
	return "rnj1_out"
}

func (p Rnj1Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "statusStg2").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926100100, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
