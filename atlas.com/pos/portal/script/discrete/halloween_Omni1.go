package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type HalloweenOmni1 struct {
}

func (p HalloweenOmni1) Name() string {
	return "halloween_Omni1"
}

func (p HalloweenOmni1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.SendPinkNotice(l, c)("SEEMS_TO_BE_LOCKED")
	return true
}
