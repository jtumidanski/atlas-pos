package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AriantAgit struct {
}

func (p AriantAgit) Name() string {
	return "ariant_Agit"
}

func (p AriantAgit) Enter(l logrus.FieldLogger, c script.Context) bool {

	if script.QuestCompleted(l, c)(3928) && script.QuestCompleted(l, c)(3931) && script.QuestCompleted(l, c)(3934) {
		character.PlayPortalSound(l)
		character.WarpById(l, c)(260000201, 1)
		return true
	} else {
		character.SendPinkNotice(l, c)("SAND_BANDITS_ONLY")
		return false
	}
}
