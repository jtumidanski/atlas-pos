package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Advice03 struct {
}

func (p Advice03) Name() string {
	return "advice03"
}

func (p Advice03) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("Press #e#b[Up]#k on the arrow key#n to climb up the ladder or rope.", 350, 5)
	return true
}
