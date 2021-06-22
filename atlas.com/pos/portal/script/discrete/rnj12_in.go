package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Rnj12In struct {
}

func (p Rnj12In) Name() string {
	return "rnj12_in"
}

func (p Rnj12In) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(926100401, 0)
	return true
}
