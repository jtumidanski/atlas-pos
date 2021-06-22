package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MapleMarket7Out struct {
}

func (p MapleMarket7Out) Name() string {
	return "mapleMarket7_out"
}

func (p MapleMarket7Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(script.GetSavedLocation(l, c)("EVENT"))
	return true
}
