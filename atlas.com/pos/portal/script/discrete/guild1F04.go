package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Guild1F04 struct {
}

func (p Guild1F04) Name() string {
	return "guild1F04"
}

func (p Guild1F04) Enter(l logrus.FieldLogger, c script.Context) bool {
	//pi.getEventInstance().gridInsert(pi.getPlayer(), 2)
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(990000700, "st00")
	return true
}
