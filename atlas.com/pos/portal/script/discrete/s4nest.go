package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4Nest struct {
}

func (p S4Nest) Name() string {
	return "s4nest"
}

func (p S4Nest) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(6241) && !script.QuestStarted(l, c)(6243) {
		script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 924000100) != 0 {
		script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}

	_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 924000100)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(924000100, 0)
	return true
}
