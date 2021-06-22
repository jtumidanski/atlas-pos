package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutPerrion2 struct {
}

func (p OutPerrion2) Name() string {
	return "outPerrion_2"
}

func (p OutPerrion2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(105100000, 0)
	return true
}
