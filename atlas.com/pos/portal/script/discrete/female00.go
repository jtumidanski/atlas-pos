package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Female00 struct {
}

func (p Female00) Name() string {
	return "female00"
}

func (p Female00) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.GetGender(l)(c.CharacterId()) == character.GenderFemale {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(670010200, 4)
		return true
	}
	script.SendPinkNotice(l, c)("CANNOT_PROCEED")
	return false
}
