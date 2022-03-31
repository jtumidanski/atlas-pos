package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Nest struct {
}

func (p S4Nest) Name() string {
	return "s4nest"
}

func (p S4Nest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(6241) && !processor.QuestStarted(l, c)(6243) {
		processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 924000100) != 0 {
		processor.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 924000100)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(924000100, 0)
	return true
}
