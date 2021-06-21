package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterTd struct {
}

func (p EnterTd) Name() string {
	return "enter_td"
}

func (p EnterTd) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(600000000, "yn00")
	return true
}
