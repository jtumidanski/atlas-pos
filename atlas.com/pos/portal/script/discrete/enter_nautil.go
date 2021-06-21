package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterNautil struct {
}

func (p EnterNautil) Name() string {
	return "enter_nautil"
}

func (p EnterNautil) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(120010000, "nt01")
	return true
}
