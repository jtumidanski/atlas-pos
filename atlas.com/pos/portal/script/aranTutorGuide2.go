package script

import (
	"log"
)

type aranTutorGuide2 struct {
}

func AranTutorGuide2() PortalScript {
	return aranTutorGuide2{}
}

func (a aranTutorGuide2) Name() string {
	return "aranTutorGuide2"
}

func (a aranTutorGuide2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002,  "cmd=o") {
		return false
	}
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide3")
	p.SendPinkNotice("ARAN_TUTORIAL_COMMAND")
	p.UpdateAreaInfo(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
