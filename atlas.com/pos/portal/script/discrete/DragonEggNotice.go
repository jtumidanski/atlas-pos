package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DragonEggNotice struct {
}

func (p DragonEggNotice) Name() string {
	return "DrnEggNotice"
}

func (p DragonEggNotice) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22014, "egg=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22014, "egg=o;mo30=o;mo40=o;mo41=o;mo50=o;mo42=o;mo60=o")
	script.ShowInfo(l, c)("UI/tutorial/evan/8/0")
	return true
}
