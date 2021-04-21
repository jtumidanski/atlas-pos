package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorArrow0 struct {
}

func AranTutorArrow0() PortalScript {
	return aranTutorArrow0{}
}

func (a aranTutorArrow0) Name() string {
	return "aranTutorArrow0"
}

func (a aranTutorArrow0) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "arr0=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "arr0=o;mo1=o;mo2=o;mo3=o")
	p.ShowInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow3")
	return true
}
