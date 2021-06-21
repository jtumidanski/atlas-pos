package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterWitch struct {
}

func (p EnterWitch) Name() string {
	return "enterWitch"
}

func (p EnterWitch) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestCompleted(l, c)(20404) {
		script.SendPinkNotice(l, c)("SHOULD_NOT_GO_CREEPY")
		return false
	}

	mapId := uint32(924010000)
	if script.QuestCompleted(l, c)(20407) {
		mapId = 924010200
	} else if script.QuestCompleted(l, c)(20406) {
		mapId = 924010100
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(mapId, 1)
	return true
}
