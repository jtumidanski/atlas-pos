package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MCRevive6 struct {
}

func (p MCRevive6) Name() string {
	return "MCRevive6"
}

func (p MCRevive6) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(980000601, 0)
	return true
}
