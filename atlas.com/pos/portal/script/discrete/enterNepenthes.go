package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterNepenthes struct {
}

func (p EnterNepenthes) Name() string {
	return "enterNepenthes"
}

func (p EnterNepenthes) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestActive(l, c)(21739) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(200060001, 2)
		return true
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 920030000) > 0 || _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 920030001) > 0 {
		script.SendPinkNotice(l, c)("SOMEONE_ALREADY_CHALLENGING")
		return false
	}

	_map.ResetPartyQuestLevel(l)(c.WorldId(), c.ChannelId(), 920030000, 1)
	_map.ResetPartyQuestLevel(l)(c.WorldId(), c.ChannelId(), 920030001, 1)

	_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 920030001, 9300348, 591, -34)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920030000, 2)
	return true
}
