package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ContactDragon struct {
}

func (p ContactDragon) Name() string {
	return "contactDragon"
}

func (p ContactDragon) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(900090100, 0)
	return true
}
