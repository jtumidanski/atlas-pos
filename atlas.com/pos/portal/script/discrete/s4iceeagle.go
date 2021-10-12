package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4IceEagle struct {
}

func (p S4IceEagle) Name() string {
	return "s4iceeagle"
}

func (p S4IceEagle) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestStarted(l, c)(6242) {
		script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 921100210) != 0 {
		script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 921100210)
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(921100210, 0)
	return true
}
