package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice02 struct {
}

func (p Advice02) Name() string {
	return "advice02"
}

func (p Advice02) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("Press #e#b[Alt]#k#n to\r\\ JUMP.", 100, 5)
	return true
}
