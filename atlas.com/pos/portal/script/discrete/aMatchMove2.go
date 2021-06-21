package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AMatchMove2 struct {
}

func (p AMatchMove2) Name() string {
	return "aMatchMove2"
}

func (p AMatchMove2) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(script.GetSavedLocation(l, c)("MIRROR"))
	return true
}
