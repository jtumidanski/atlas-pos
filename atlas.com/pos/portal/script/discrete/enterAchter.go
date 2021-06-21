package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterAchter struct {
}

func (p EnterAchter) Name() string {
	return "enterAchter"
}

func (p EnterAchter) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpByName(l, c)(100000201, "out02")
	return true
}
