package script

import (
	"log"
)

type aranTutorMono0 struct {
}

func AranTutorMono0() PortalScript {
	return aranTutorMono0{}
}

func (a aranTutorMono0) Name() string {
	return "aranTutorMono0"
}

func (a aranTutorMono0) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "mo1=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "mo1=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon1")
	return true
}
