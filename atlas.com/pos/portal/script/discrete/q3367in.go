package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3367In struct {
}

func (p Q3367In) Name() string {
	return "q3367in"
}

func (p Q3367In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(3367) {
		processor.SendPinkNotice(l, c)("ROOM_NO_ACCESS")
		return false
	}
	bd := uint32(processor.QuestProgressInt(l, c)(3367, 31))
	bi := character.ItemQuantity(l)(c.CharacterId(), 4031797)
	if bi < bd {
		processor.GainItem(l, c)(4031797, int16(bd-bi))
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(926130102, 0)
	return true
}
