package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice09 struct {
}

func (p Advice09) Name() string {
	return "advice09"
}

func (p Advice09) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("Press #e#b[Down]#k on the arrow key#n and#e#b[Alt]#k#n at the same time to jump downwards.", 450, 6)
	return true
}
