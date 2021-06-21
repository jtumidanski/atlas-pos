package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Dracoout struct {
}

func (p Dracoout) Name() string {
	return "dracoout"
}

func (p Dracoout) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240000100, "east00")
	return true
}
