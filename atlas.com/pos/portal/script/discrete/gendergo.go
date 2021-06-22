package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Gendergo struct {
}

func (p Gendergo) Name() string {
	return "gende"
}

func (p Gendergo) Enter(l logrus.FieldLogger, c script.Context) bool {
	if c.Portal().Name() == "female00" {
		if character.GetGender(l)(c.CharacterId()) == character.GenderFemale {
			script.PlayPortalSound(l, c)
			script.WarpByName(l, c)(c.MapId(), "female01")
			return true
		}
		script.SendPinkNotice(l, c)("FEMALE_ONLY")
		return false
	}
	if character.GetGender(l)(c.CharacterId()) == character.GenderMale {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(c.MapId(), "male01")
		return true
	}
	script.SendPinkNotice(l, c)("MALE_ONLY")
	return false
}
