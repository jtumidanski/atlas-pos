package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal7 struct {
}

func (p GlpqPortal7) Name() string {
	return "glpqPortal7"
}

func (p GlpqPortal7) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(610030800, 0)
	return true
}
