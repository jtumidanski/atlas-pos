package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CannonTuto10 struct {
}

func (p CannonTuto10) Name() string {
	return "cannon_tuto_10"
}

func (p CannonTuto10) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.UnlockUI(l, c)
	return true
}
