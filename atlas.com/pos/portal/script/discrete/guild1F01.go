package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Guild1F01 struct {
}

func (p Guild1F01) Name() string {
	return "guild1F01"
}

func (p Guild1F01) Enter(l logrus.FieldLogger, c script.Context) bool {
	//pi.getEventInstance().gridInsert(pi.getPlayer(), 0)
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(990000700, "st00")
	return true
}
