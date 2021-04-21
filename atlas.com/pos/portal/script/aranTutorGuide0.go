package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorGuide0 struct {
}

func AranTutorGuide0() PortalScript {
	return aranTutorGuide0{}
}

func (a aranTutorGuide0) Name() string {
	return "aranTutorGuide0"
}

func (a aranTutorGuide0) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "normal=o") {
		return false
	}
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide1")
	p.SendPinkNotice("ARAN_TUTORIAL_REGULAR_ATTACK")
	p.UpdateAreaInfo(21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o")
	return true
}
