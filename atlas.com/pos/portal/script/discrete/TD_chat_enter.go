package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDChatEnter struct {
}

func (p TDChatEnter) Name() string {
	return "TD_chat_enter"
}

func (p TDChatEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.OpenNPC(l, c)(2083006)
	return false
}
