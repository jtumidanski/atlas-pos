package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlTutoMsg0 struct {
}

func (p GlTutoMsg0) Name() string {
	return "glTutoMsg0"
}

func (p GlTutoMsg0) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.ShowInstruction(l, c)("Once you leave this area you won't be able to return.", 150, 5)
	return true
}
