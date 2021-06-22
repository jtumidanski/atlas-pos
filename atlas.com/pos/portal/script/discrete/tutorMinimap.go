package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TutorMiniMap struct {
}

func (p TutorMiniMap) Name() string {
	return "tutorMinimap"
}

func (p TutorMiniMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.GuideHint(l, c)(1)
	script.BlockPortal(l, c)
	return true
}
