package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InERShip struct {
}

func (p InERShip) Name() string {
	return "inERShip"
}

func (p InERShip) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(101000400, 2)
	return true
}
