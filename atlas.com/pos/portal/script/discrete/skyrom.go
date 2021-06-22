package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Skyrom struct {
}

func (p Skyrom) Name() string {
	return "skyrom"
}

func (p Skyrom) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(3935) || script.HasItem(l, c)(4031574) {
		return false
	}
	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 926000010) != 0 {
		script.SendPinkNotice(l, c)("OTHER_PLAYER_TRYING")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(926000010, 0)
	return true
}
