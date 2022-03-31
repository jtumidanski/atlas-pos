package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDChatEnter struct {
}

func (p TDChatEnter) Name() string {
	return "TD_chat_enter"
}

func (p TDChatEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPC(l, c)(2083006)
	return false
}
