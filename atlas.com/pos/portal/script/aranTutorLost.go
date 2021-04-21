package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorLost struct {
}

func AranTutorLost() PortalScript {
	return aranTutorLost{}
}

func (a aranTutorLost) Name() string {
	return "aranTutorLost"
}

func (a aranTutorLost) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.BlockPortal()
	if p.ContainsAreaInfo(21002, "fin=o") {
		return false
	}
	p.UpdateAreaInfo(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;fin=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	p.ShowIntro("Effect/Direction1.img/aranTutorial/ClickChild")
	return true
}
