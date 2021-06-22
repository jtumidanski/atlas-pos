package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type HorntaleMorph2 struct {
}

func (p HorntaleMorph2) Name() string {
	return "hontale_morph2"
}

func (p HorntaleMorph2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(240040600, 4)
	return true
}
