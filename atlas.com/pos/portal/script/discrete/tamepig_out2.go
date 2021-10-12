package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TamePigOut2 struct {
}

func (p TamePigOut2) Name() string {
	return "tamepig_out2"
}

func (p TamePigOut2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !(script.HasItems(l, c)(4031507, 5) && script.HasItems(l, c)(4031508, 5) && script.QuestStarted(l, c)(6002)) {
		character.RemoveAll(l)(c.CharacterId(), 4031507)
		character.RemoveAll(l)(c.CharacterId(), 4031508)
	}
	pc := character.ItemQuantity(l)(c.CharacterId(), 4031507)
	rc := character.ItemQuantity(l)(c.CharacterId(), 4031508)

	if pc > 5 {
		script.GainItem(l, c)(4031507, -1 * (int16(pc) - 5))
	}
	if rc > 5 {
		script.GainItem(l, c)(4031508, -1 * (int16(rc) - 5))
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(230000003, "out00")
	return true
}
