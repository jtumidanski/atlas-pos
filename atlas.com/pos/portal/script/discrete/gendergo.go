package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Gendergo struct {
}

func (p Gendergo) Name() string {
	return "gende"
}

func (p Gendergo) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.Portal().Name() == "female00" {
		if character.GetGender(l, span)(c.CharacterId()) == character.GenderFemale {
			processor.PlayPortalSound(l, c)
			processor.WarpByName(l, span, c)(c.MapId(), "female01")
			return true
		}
		processor.SendPinkNotice(l, c)("FEMALE_ONLY")
		return false
	}
	if character.GetGender(l, span)(c.CharacterId()) == character.GenderMale {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(c.MapId(), "male01")
		return true
	}
	processor.SendPinkNotice(l, c)("MALE_ONLY")
	return false
}
