package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type JnrExit struct {
}

func (p JnrExit) Name() string {
	return "jnr_exit"
}

func (p JnrExit) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(261000021, 0)
	return true
}
