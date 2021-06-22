package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutPerrion1 struct {
}

func (p OutPerrion1) Name() string {
	return "outPerrion_1"
}

func (p OutPerrion1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.SendPinkNotice(l, c)("TEMPLE_SHORTCUT")
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(105100000, 2)
	return true
}
