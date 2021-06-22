package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MinrarJob4 struct {
}

func (p MinrarJob4) Name() string {
	return "minar_job4"
}

func (p MinrarJob4) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240010501, "out00")
	return true
}
