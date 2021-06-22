package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GagaSuccess struct {
}

func (p GagaSuccess) Name() string {
	return "gaga_success"
}

func (p GagaSuccess) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(922240100 + (c.MapId() - 922240000), 0)
	return true
}
