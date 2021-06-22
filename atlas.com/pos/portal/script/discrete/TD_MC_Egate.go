package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TDMCEGate struct {
}

func (p TDMCEGate) Name() string {
	return "TD_MC_Egate"
}

func (p TDMCEGate) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(106021300, 1)
	return true
}
