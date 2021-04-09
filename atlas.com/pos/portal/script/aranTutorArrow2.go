package script

import (
	"log"
)

type aranTutorArrow2 struct {
}

func AranTutorArrow2() PortalScript {
	return aranTutorArrow2{}
}

func (a aranTutorArrow2) Name() string {
	return "aranTutorArrow2"
}

func (a aranTutorArrow2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "arr2=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "normal=o;arr0=o;arr1=o;arr2=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
