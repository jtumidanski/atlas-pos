package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutTestWolf struct {
}

func (p OutTestWolf) Name() string {
	return "outtestWolf"
}

func (p OutTestWolf) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) != 0 {
		processor.SendPinkNotice(l, c)("DEFEAT_ALL_WOLVES")
		return false
	}
	if !processor.CanHold(l, c)(4001193, 1) {
		processor.SendPinkNotice(l, c)("FREE_SPACE_FOR_COUSE_CLEAR_TOKEN")
		return false
	}
	processor.GainItem(l, c)(4001193, 1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(140010210, 0)
	return true
}
