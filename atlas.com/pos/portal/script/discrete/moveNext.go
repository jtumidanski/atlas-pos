package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MoveNext struct {
}

func (p MoveNext) Name() string {
	return "moveNext"
}

func (p MoveNext) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(c.MapId() + 10, "east00")
	return true
}
