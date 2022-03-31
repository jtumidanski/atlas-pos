package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party6Stage struct {
}

func (p Party6Stage) Name() string {
	return "party6_stage"
}

func (p Party6Stage) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	switch c.MapId() {
	case 930000000:
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(930000100, 0)
		return true
	case 930000100:
		if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 0 {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(930000200, 0)
			return true
		} else {
			processor.SendPinkNotice(l, c)("ELIMINATE_MONSTERS")
			return false
		}
	case 930000200:
		if r, err := reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), "spine"); err == nil && r.State() < 4 {
			processor.SendPinkNotice(l, c)("SPINE_BLOCKS")
			return false
		} else {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(930000300, 0) //assuming they cant get past reactor without it being gone
			return true
		}
	default:
		processor.SendPinkNotice(l, c)("UNBOUND_PATH")
		return false
	}
}
