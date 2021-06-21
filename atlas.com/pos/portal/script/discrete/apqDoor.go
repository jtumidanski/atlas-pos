package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ApqDoor struct {
}

func (p ApqDoor) Name() string {
	return "apqDoor"
}

func (p ApqDoor) Enter(l logrus.FieldLogger, c script.Context) bool {
	r, err := script.GetReactor(l, c)(c.MapId(), "gate"+c.Portal().Name())
	if err != nil || r.State() != 4 {
		script.SendPinkNotice(l, c)("GATE_NOT_YET_OPEN")
		return false
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(670010600, "gt"+c.Portal().Name()+"PIB")
	return true
}
