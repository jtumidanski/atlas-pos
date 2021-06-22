package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MCRevive5 struct {
}

func (p MCRevive5) Name() string {
	return "MCRevive5"
}

func (p MCRevive5) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(980000501, 0)
	return true
}
