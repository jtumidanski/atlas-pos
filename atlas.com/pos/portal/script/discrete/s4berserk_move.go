package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4BerserkMove struct {
}

func (p S4BerserkMove) Name() string {
	return "s4berserk_move"
}

func (p S4BerserkMove) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 0 {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(910500200, "out00")
		return true
	}
	processor.SendPinkNotice(l, c)("DEFEAT_ALL_MONSTERS_FIRST")
	return true
}
