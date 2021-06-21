package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ApqClosed struct {
}

func (p ApqClosed) Name() string {
	return "apqClosed"
}

func (p ApqClosed) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.SendPinkNotice(l, c)("GATE_IS_NOT_YET_OPENED")
	return false
}
