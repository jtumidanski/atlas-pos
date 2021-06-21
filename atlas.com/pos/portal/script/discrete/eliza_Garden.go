package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ElizaGarden struct {
}

func (p ElizaGarden) Name() string {
	return "eliza_Garden"
}

func (p ElizaGarden) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920020000, 2)
	return true
}
