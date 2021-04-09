package script

import (
	"log"
)

type aranTutorMono1 struct {
}

func AranTutorMono1() PortalScript {
	return aranTutorMono1{}
}

func (a aranTutorMono1) Name() string {
	return "aranTutorMono1"
}

func (a aranTutorMono1) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "mo2=o") {
		return false
	}
	p.PlaySound("Aran/balloon")
	p.UpdateAreaInfo( 21002, "mo1=o;mo2=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon2")
	return true
}
