package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
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
	if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "statusStg3"); err == nil && r.State() == 3 {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(926100200, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
