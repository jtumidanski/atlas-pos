package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GuildWaitingExit struct {
}

func (p GuildWaitingExit) Name() string {
	return "guildwaitingexit"
}

func (p GuildWaitingExit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpRandom(l, span, c)(101030104)
	return true
}
