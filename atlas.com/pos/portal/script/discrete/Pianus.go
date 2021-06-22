package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Pianus struct {
}

func (p Pianus) Name() string {
	return "Pianus"
}

func (p Pianus) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(230040420, "out00")
	return true
}
