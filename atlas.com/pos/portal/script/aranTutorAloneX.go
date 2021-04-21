package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorAloneX struct {
}

func AranTutorAloneX() PortalScript {
	return aranTutorAloneX{}
}

func (a aranTutorAloneX) Name() string {
	return "aranTutorAloneX"
}

func (a aranTutorAloneX) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(914000100, 1)
	return true
}
