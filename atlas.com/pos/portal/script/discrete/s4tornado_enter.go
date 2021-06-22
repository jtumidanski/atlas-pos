package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4TornadoEnter struct {
}

func (p S4TornadoEnter) Name() string {
	return "s4tornado_enter"
}

func (p S4TornadoEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(6230) || script.QuestStarted(l, c)(6231) || script.HasItem(l, c)(4001110) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 922020200) == 0 {
			_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 922020200)
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(922020200, 0)
			return true
		}
		script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
		return false
	}
	return false
}
