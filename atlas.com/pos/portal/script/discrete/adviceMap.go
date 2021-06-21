package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AdviceMap struct {
}

func (p AdviceMap) Name() string {
	return "adviceMap"
}

func (p AdviceMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("Press the #e#b[Up]#k arrow#n to use the portal and move to the next map.", 230, 5)
	return true
}
