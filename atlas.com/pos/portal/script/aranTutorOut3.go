package script

import (
	"github.com/sirupsen/logrus"
)

type aranTutorOut3 struct {
}

func AranTutorOut3() PortalScript {
	return aranTutorOut3{}
}

func (a aranTutorOut3) Name() string {
	return "aranTutorOut3"
}

func (a aranTutorOut3) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.TeachSkill(20000016, 0, -1, -1)
	p.TeachSkill(20000016, 1, 0, -1)
	p.PlayPortalSound()
	p.WarpById(914000220, 1)
	return true
}
