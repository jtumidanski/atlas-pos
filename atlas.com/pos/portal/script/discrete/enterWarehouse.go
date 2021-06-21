package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterWarehouse struct {
}

func (p EnterWarehouse) Name() string {
	return "enterWarehouse"
}

func (p EnterWarehouse) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(300000011, 0)
	return true
}
