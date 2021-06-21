package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CatPriestMap struct {
}

func (p CatPriestMap) Name() string {
	return "catPriest_map"
}

func (p CatPriestMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(925000000, 2)
	return true
}
