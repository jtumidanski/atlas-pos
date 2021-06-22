package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutMaha struct {
}

func (p OutMaha) Name() string {
	return "outMaha"
}

func (p OutMaha) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(140000000, 0)
	return true
}
