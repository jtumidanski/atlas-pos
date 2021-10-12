package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MinrarElli struct {
}

func (p MinrarElli) Name() string {
	return "minar_elli"
}

func (p MinrarElli) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(4031346) {
		script.SendLightBlueNotice(l, c)("MAGIC_SEED_NEEDED")
		return false
	}
	if c.MapId() == 240010100 {
		script.GainItem(l, c)(4031346, -1)
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(101010000, "minar00")
		return true
	}
	if c.MapId() == 101010000 {
		script.GainItem(l, c)(4031346, -1)
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(240010100, "elli00")
		return true
	}
	return true
}
