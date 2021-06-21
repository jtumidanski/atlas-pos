package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice07 struct {
}

func (p Advice07) Name() string {
	return "advice07"
}

func (p Advice07) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("You can view the World Map by pressing the #e#b[W]#k#nkey.", 350, 5)
	return true
}
