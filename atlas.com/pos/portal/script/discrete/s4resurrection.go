package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4Resurrection struct {
}

func (p S4Resurrection) Name() string {
	return "s4resurrection"
}

func (p S4Resurrection) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(4001108) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 923000100) == 0 {
			_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 923000100)
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(923000100, 0)
			return true
		} else {
			script.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
			return false
		}
	}
	script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
