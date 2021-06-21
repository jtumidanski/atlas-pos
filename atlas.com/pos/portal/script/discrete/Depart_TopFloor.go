package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DepartTopFloor struct {
}

func (p DepartTopFloor) Name() string {
	return "Depart_TopFloor"
}

func (p DepartTopFloor) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.OpenNPC(l, c)(1052125)
	return true
}
