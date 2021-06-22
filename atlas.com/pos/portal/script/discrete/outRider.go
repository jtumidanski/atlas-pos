package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutRider struct {
}

func (p OutRider) Name() string {
	return "outRider"
}

func (p OutRider) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.CanHold(l, c)(4001193, 1) {
		script.SendPinkNotice(l, c)("FREE_SPACE_FOR_COUSE_CLEAR_TOKEN")
		return false
	}
	script.GainItem(l, c)(4001193, 1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(211050000, 4)
	return true
}
