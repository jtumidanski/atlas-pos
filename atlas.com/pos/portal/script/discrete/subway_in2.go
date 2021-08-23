package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type SubwayIn2 struct {
}

func (p SubwayIn2) Name() string {
	return "subway_in2"
}

func (p SubwayIn2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(103000101, 3)
	return true
}