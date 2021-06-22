package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4Common1Exit struct {
}

func (p S4Common1Exit) Name() string {
	return "s4common1_exit"
}

func (p S4Common1Exit) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	if !script.HasItem(l, c)(4031495) {
		script.WarpRandom(l, c)(211040100)
		return true
	}
	script.WarpRandom(l, c)(921100301)
	return true
}
