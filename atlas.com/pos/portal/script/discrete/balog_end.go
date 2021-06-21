package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type BalogEnd struct {
}

func (p BalogEnd) Name() string {
	return "balog_end"
}

func (p BalogEnd) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.CanHold(l, c)(4001261, 1) {
		character.GainItem(l, c)(4001261, 1)
		character.PlayPortalSound(l)
		character.WarpById(l, c)(105100100, 0)
		return true
	}
	character.SendPinkNotice(l, c)("INVENTORY_FULL_ERROR")
	return false
}
