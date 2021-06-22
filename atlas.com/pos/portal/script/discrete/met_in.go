package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MetIn struct {
}

func (p MetIn) Name() string {
	return "met_in"
}

func (p MetIn) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(103000103, 1)
	return true
}
