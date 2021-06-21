package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AriantQueens struct {
}

func (p AriantQueens) Name() string {
	return "ariant_queens"
}

func (p AriantQueens) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.Morphed(l, c)(221005) {
		return false
	} else {
		character.PlayPortalSound(l)
		character.WarpById(l, c)(260000300, 7)
		character.SendPinkNotice(l, c)("PALACE_INTRUDER")
		return true
	}
}
