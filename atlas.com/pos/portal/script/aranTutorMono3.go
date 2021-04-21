package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorMono3 struct {
}

func AranTutorMono3() PortalScript {
	return aranTutorMono3{}
}

func (a aranTutorMono3) Name() string {
	return "aranTutorMono3"
}

func (a aranTutorMono3) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "mo4=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o;mo4=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon6")
	return true
}
