package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanEntrance struct {
}

func (p EvanEntrance) Name() string {
	return "evanEntrance"
}

func (p EvanEntrance) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(100030400, "east00")
	return true
}
