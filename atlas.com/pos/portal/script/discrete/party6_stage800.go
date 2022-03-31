package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party6Stage800 struct {
}

func (p Party6Stage800) Name() string {
	return "party6_stage800"
}

func (p Party6Stage800) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	character.RemoveAll(l)(c.CharacterId(), 4001162)
	character.RemoveAll(l)(c.CharacterId(), 4001163)
	character.RemoveAll(l)(c.CharacterId(), 4001164)
	character.RemoveAll(l)(c.CharacterId(), 4001169)
	character.RemoveAll(l)(c.CharacterId(), 2270004)

	r, err := reactor.GetById(l, span)(3008000)
	if err == nil && r.State() > 0 {
		if !processor.CanHold(l, c)(4001198, 1) {
			processor.SendPinkNotice(l, c)("INVENTORY_FULL")
			return false
		}
		processor.GainItem(l, c)(4001198, 1)
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(300030100, 0)
	return true
}
