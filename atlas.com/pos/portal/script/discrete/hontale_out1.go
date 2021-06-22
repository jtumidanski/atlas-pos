package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type HorntaleOut1 struct {
}

func (p HorntaleOut1) Name() string {
	return "hontale_out1"
}

func (p HorntaleOut1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240050400, "sp")
	return true
}
