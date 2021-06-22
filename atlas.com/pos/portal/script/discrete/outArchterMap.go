package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutArchterMap struct {
}

func (p OutArchterMap) Name() string {
	return "outArchterMap"
}

func (p OutArchterMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(100000000, "Achter00")
	return true
}
