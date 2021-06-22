package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type SpaceReturn struct {
}

func (p SpaceReturn) Name() string {
	return "space_return"
}

func (p SpaceReturn) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(script.GetSavedLocation(l, c)("EVENT"))
	return true
}
