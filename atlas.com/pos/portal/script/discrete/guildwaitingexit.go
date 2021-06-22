package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GuildWaitingExit struct {
}

func (p GuildWaitingExit) Name() string {
	return "guildwaitingexit"
}

func (p GuildWaitingExit) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(101030104)
	return true
}
