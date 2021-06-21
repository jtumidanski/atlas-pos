package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterFirstDH struct {
}

func (p EnterFirstDH) Name() string {
	return "enterFirstDH"
}

func (p EnterFirstDH) Enter(l logrus.FieldLogger, c script.Context) bool {
	mapId := uint32(0)
	if script.QuestStarted(l, c)(20701) {
		mapId = 913000000
	} else if script.QuestStarted(l, c)(20702) {
		mapId = 913000100
	} else if script.QuestStarted(l, c)(20703) {
		mapId = 913000200
	}

	if mapId == 0 {
		script.SendPinkNotice(l, c)("KIKUS_ACCLIMATION_TRAINING_REQUIREMENT")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), mapId) > 0 {
		script.SendPinkNotice(l, c)("SOMEONE_ALREADY_IN_MAP")
		return false
	}

	_map.ResetPartyQuest(l)(c.WorldId(), c.ChannelId(), mapId)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(mapId, 0)
	return true
}
