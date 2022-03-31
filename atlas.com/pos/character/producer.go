package character

import (
	"atlas-pos/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

func emitEnableActions(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_ENABLE_ACTIONS")
	return func(characterId uint32) {
		event := &enableActionsEvent{characterId}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}
