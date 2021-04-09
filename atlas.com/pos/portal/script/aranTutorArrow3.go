package script

import (
	"log"
)

type aranTutorArrow3 struct {
}

func AranTutorArrow3() PortalScript {
	return aranTutorArrow3{}
}

func (a aranTutorArrow3) Name() string {
	return "aranTutorArrow3"
}

func (a aranTutorArrow3) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "arr3=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
