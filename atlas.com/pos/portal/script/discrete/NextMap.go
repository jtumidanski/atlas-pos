package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type NextMap struct {
}

func (p NextMap) Name() string {
	return "NextMap"
}

func (p NextMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(c.MapId() + 100, 0)
	return true
}
