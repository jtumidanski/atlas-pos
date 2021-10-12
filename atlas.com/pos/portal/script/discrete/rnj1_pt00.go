package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj1Pt00 struct {
}

func (p Rnj1Pt00) Name() string {
	return "rnj1_pt00"
}

func (p Rnj1Pt00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "statusStg1").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(926100001, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
