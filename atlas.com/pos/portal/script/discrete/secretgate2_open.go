package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type SecretGate2Open struct {
}

func (p SecretGate2Open) Name() string {
	return "secretgate2_open"
}

func (p SecretGate2Open) Enter(l logrus.FieldLogger, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "secretgate2").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(990000631,1)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_IS_CLOSED")
	return false
}
