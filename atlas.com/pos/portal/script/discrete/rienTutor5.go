package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RienTutor5 struct {
}

func (p RienTutor5) Name() string {
	return "rienTutor5"
}

func (p RienTutor5) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.TalkGuide(l, c)("You're very close to town. I'll head over there first since I have some things to take care of. You take your time.")
	script.BlockPortal(l, c)
	return false
}
