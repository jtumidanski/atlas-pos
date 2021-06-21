package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DepartTopOut struct {
}

func (p DepartTopOut) Name() string {
	return "Depart_topOut"
}

func (p DepartTopOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(103040300, 1)
	return true
}
