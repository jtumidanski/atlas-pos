package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterRienFirst struct {
}

func (p EnterRienFirst) Name() string {
	return "enterRienFirst"
}

func (p EnterRienFirst) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.IsJob(l, span)(c.CharacterId(), job.Legend) &&
		!processor.QuestCompleted(l, c)(21014) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(140000000, "st00")
	} else {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(140000000, "west00")
	}
	return true
}
