package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice01 struct {
}

func (p Advice01) Name() string {
	return "advice01"
}

func (p Advice01) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("Click \r\\#b<Heena>#k", 100, 5)
	return true
}
