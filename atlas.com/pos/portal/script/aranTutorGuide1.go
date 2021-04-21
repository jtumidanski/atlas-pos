package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorGuide1 struct {
}

func AranTutorGuide1() PortalScript {
	return aranTutorGuide1{}
}

func (a aranTutorGuide1) Name() string {
	return "aranTutorGuide1"
}

func (a aranTutorGuide1) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "chain=o") {
		return false
	}
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide2")
	p.SendPinkNotice("ARAN_TUTORIAL_CONSECUTIVE")
	p.UpdateAreaInfo(21002, "normal=o;arr0=o;arr1=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
