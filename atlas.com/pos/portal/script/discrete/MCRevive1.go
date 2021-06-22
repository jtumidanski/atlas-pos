package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MCRevive1 struct {
}

func (p MCRevive1) Name() string {
	return "MCRevive1"
}

func (p MCRevive1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(980000101, 0)
	return true
}
