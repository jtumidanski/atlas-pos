package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice06 struct {
}

func (p Advice06) Name() string {
	return "advice06"
}

func (p Advice06) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("Press the #e#b[Up]#k arrow#n to use the portal \r\\and move to the next map.", 230, 5)
	return true
}
