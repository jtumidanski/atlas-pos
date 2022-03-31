package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4FireHawk struct {
}

func (p S4FireHawk) Name() string {
	return "s4firehawk"
}

func (p S4FireHawk) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(6240) {
		processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 921100200) != 0 {
		processor.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 921100200)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(921100200, 0)
	return true
}
