package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanGarden0 struct {
}

func (p EvanGarden0) Name() string {
	return "evanGarden0"
}

func (p EvanGarden0) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(100030200, "east00")
	return true
}
