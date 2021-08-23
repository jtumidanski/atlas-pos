package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type Party6Stage512 struct {
}

func (p Party6Stage512) Name() string {
	return "party6_stage512"
}

func (p Party6Stage512) Enter(l logrus.FieldLogger, c script.Context) bool {
	if rand.Float64() < 0.1 {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(930000300, "16st")
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(930000300, "13st")
	return true
}