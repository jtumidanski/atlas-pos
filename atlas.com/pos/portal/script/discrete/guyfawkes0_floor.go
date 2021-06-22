package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GuyFawkes0Floor struct {
}

func (p GuyFawkes0Floor) Name() string {
	return "guyfawkes0_floor"
}

func (p GuyFawkes0Floor) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.WarpById(l, c)(674030000, 0)
	return true
}
