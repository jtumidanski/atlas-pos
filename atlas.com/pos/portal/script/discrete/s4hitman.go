package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Hitman struct {
}

func (p S4Hitman) Name() string {
	return "s4hitman"
}

func (p S4Hitman) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(6201) {
		processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 910200000) != 0 {
		processor.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 910200000)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(910200000, 0)
	return true
}
