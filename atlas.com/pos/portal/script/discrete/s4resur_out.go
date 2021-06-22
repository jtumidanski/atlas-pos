package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4ResurOut struct {
}

func (p S4ResurOut) Name() string {
	return "s4resur_out"
}

func (p S4ResurOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(6134) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(220070400, 3)
		return true
	}
	if !script.CanHold(l, c)(4031448, 1) {
		script.SendPinkNotice(l, c)("MAKE_ROOM_FOR_QUEST_ITEM")
		return false
	}
	script.GainItem(l, c)(4031448, 1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(220070400, 3)
	return true
}
