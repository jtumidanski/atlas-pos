package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3367In struct {
}

func (p Q3367In) Name() string {
	return "q3367in"
}

func (p Q3367In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestStarted(l, c)(3367) {
		script.SendPinkNotice(l, c)("ROOM_NO_ACCESS")
		return false
	}
	bd := uint32(script.QuestProgressInt(l, c)(3367, 31))
	bi := character.ItemQuantity(l)(c.CharacterId(), 4031797)
	if bi < bd {
		script.GainItem(l, c)(4031797, int16(bd-bi))
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(926130102, 0)
	return true
}
