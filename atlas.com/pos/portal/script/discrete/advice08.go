package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice08 struct {
}

func (p Advice08) Name() string {
	return "advice08"
}

func (p Advice08) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("You can check your character's stats by pressing the #e#b[S]#k#nkey.", 350, 5)
	return true
}
