package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ChristmasOut2008 struct {
}

func (p ChristmasOut2008) Name() string {
	return "08_xmas_out"
}

func (p ChristmasOut2008) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(c.MapId()-2, 0)
	return true
}
