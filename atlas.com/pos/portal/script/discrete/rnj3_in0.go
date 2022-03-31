package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj3In0 struct {
}

func (p Rnj3In0) Name() string {
	return "rnj3_in0"
}

func (p Rnj3In0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "rnj3_out1").State() == 1 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926100201, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
