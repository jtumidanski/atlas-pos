package script

import (
	"log"
)

type aranTutorArrow1 struct {
}

func AranTutorArrow1() PortalScript {
	return aranTutorArrow1{}
}

func (a aranTutorArrow1) Name() string {
	return "aranTutorArrow1"
}

func (a aranTutorArrow1) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "arr1=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "normal=o;arr0=o;arr1=o;mo1=o;mo2=o;mo3=o;mo4=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
