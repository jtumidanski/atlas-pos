package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Berserk struct {
}

func (p S4Berserk) Name() string {
	return "s4berserk"
}

func (p S4Berserk) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(6153) && script.HasItem(l, c)(4031475) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 910500200) == 0 {
			_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 910500200)
			_map.ShuffleReactors(l)(c.WorldId(), c.ChannelId(), 910500200)
			script.PlayPortalSound(l, c)
			script.WarpByName(l, span, c)(910500200, "out01")
			return true
		} else {
			script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
			return false
		}
	}
	script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
