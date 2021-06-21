package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice00 struct {
}

func (p Advice00) Name() string {
	return "advice00"
}

func (p Advice00) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("You can move by using the arrow keys.", 250, 5)
	return true
}
