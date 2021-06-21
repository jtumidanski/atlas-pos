package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice05 struct {
}

func (p Advice05) Name() string {
	return "advice05"
}

func (p Advice05) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("Press #e#b[Q]#k#n to view the Quest window.", 250, 5)
	return true
}
