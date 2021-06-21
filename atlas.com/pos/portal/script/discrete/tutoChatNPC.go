package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TutoChatNPC struct {
}

func (p TutoChatNPC) Name() string {
	return "tutoChatNPC"
}

func (p TutoChatNPC) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.HasLevel30Character(l, c) {
		script.OpenNPC(l, c)(2007)
	}
	script.BlockPortal(l, c)
	return true
}
