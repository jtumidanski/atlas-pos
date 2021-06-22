package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TutorHelper struct {
}

func (p TutorHelper) Name() string {
	return "tutorHelper"
}

func (p TutorHelper) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.SpawnGuide(l, c)
	script.TalkGuide(l, c)("Welcome to Maple World! I'm Mimo. I'm in charge of guiding you until you reach Lv. 10 and become a Knight-In-Training. Double-click for further information!")
	script.BlockPortal(l, c)
	return true
}
