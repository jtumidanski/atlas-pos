package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice04 struct {
}

func (p Advice04) Name() string {
	return "advice04"
}

func (p Advice04) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.ShowInstruction(l, c)("Click \r\\#b<Sera>", 100, 5)
	return true
}
