package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type BalogTemple struct {
}

func (p BalogTemple) Name() string {
	return "balogTemple"
}

func (p BalogTemple) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(105100000, 2)
	return true
}
