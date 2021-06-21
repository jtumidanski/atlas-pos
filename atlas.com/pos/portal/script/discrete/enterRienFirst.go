package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterRienFirst struct {
}

func (p EnterRienFirst) Name() string {
	return "enterRienFirst"
}

func (p EnterRienFirst) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.IsJob(l)(c.CharacterId(), job.Legend) &&
		!script.QuestCompleted(l, c)(21014) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(140000000, "st00")
	} else {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(140000000, "west00")
	}
	return true
}
