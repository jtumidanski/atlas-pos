package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type HalloweenEnter struct {
}

func (p HalloweenEnter) Name() string {
	return "halloween_enter"
}

func (p HalloweenEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(682000100, "st00")
	return true
}
