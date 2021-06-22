package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MoveBefore struct {
}

func (p MoveBefore) Name() string {
	return "moveBefore"
}

func (p MoveBefore) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(c.MapId() - 10, "west00")
	return true
}
