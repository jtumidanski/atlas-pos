package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type In2159011 struct {
}

func (p In2159011) Name() string {
	return "in2159011"
}

func (p In2159011) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.OpenNPC(l, c)(2159011)
	return true
}
