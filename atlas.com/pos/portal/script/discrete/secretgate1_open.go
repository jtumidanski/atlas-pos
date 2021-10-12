package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type SecretGate1Open struct {
}

func (p SecretGate1Open) Name() string {
	return "secretgate1_open"
}

func (p SecretGate1Open) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "secretgate1").State() == 1 {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(990000611,1)
		return true
	}
	script.SendPinkNotice(l, c)("DOOR_IS_CLOSED")
	return false
}
