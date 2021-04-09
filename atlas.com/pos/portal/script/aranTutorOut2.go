package script

import (
	"log"
)

type aranTutorOut2 struct {
}

func AranTutorOut2() PortalScript {
	return aranTutorOut2{}
}

func (a aranTutorOut2) Name() string {
	return "aranTutorOut2"
}

func (a aranTutorOut2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.TeachSkill(20000014, 0, -1, -1)
	p.TeachSkill(20000015, 0, -1, -1)
	p.TeachSkill(20000014, 1, 0, -1)
	p.TeachSkill(20000015, 1, 0, -1)
	p.PlayPortalSound()
	p.WarpById(914000210, 1)
	return true
}
