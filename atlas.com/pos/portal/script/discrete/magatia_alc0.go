package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MagatiaAlc0 struct {
}

func (p MagatiaAlc0) Name() string {
	return "magatia_alc0"
}

func (p MagatiaAlc0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(3309) || script.HasItem(l, c)(4031708) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(261020700, "down00")
	} else {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(926120000, "out00")
	}
	return true
}
