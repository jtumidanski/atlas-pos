package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AriantMount2 struct {
}

func (p AriantMount2) Name() string {
	return "ariantMount2"
}

func (p AriantMount2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(980010000, 0)
	return true
}
