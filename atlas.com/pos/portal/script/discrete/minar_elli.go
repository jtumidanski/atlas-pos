package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MinrarElli struct {
}

func (p MinrarElli) Name() string {
	return "minar_elli"
}

func (p MinrarElli) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(4031346) {
		processor.SendLightBlueNotice(l, c)("MAGIC_SEED_NEEDED")
		return false
	}
	if c.MapId() == 240010100 {
		processor.GainItem(l, c)(4031346, -1)
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(101010000, "minar00")
		return true
	}
	if c.MapId() == 101010000 {
		processor.GainItem(l, c)(4031346, -1)
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(240010100, "elli00")
		return true
	}
	return true
}
