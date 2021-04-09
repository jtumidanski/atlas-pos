package script

import (
	"log"
)

type aranTutorMono2 struct {
}

func AranTutorMono2() PortalScript {
	return aranTutorMono2{}
}

func (a aranTutorMono2) Name() string {
	return "aranTutorMono2"
}

func (a aranTutorMono2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "mo3=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "mo1=o;mo2=o;mo3=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon3")
	return true
}
