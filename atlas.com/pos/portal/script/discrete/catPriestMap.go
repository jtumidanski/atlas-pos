package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CatPriestMap struct {
}

func (p CatPriestMap) Name() string {
	return "catPriest_map"
}

func (p CatPriestMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(925000000, 2)
	return true
}
