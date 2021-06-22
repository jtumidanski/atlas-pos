package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RienTutor6 struct {
}

func (p RienTutor6) Name() string {
	return "rienTutor6"
}

func (p RienTutor6) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.RemoveGuide(l, c)
	script.BlockPortal(l, c)
	return true
}
