package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ReundoDraco struct {
}

func (p ReundoDraco) Name() string {
	return "reundodraco"
}

func (p ReundoDraco) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	return true
}
