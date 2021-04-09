package script

import (
	"log"
)

type aranTutorOut1 struct {
}

func AranTutorOut1() PortalScript {
	return aranTutorOut1{}
}

func (a aranTutorOut1) Name() string {
	return "aranTutorOut1"
}

func (a aranTutorOut1) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	if p.QuestStarted(21000) {
		p.TeachSkill(20000017, 0, -1, -1)
		p.TeachSkill(20000018, 0, -1, -1)
		p.TeachSkill(20000018, 1, 0, -1)
		p.PlayPortalSound()
		p.WarpById(914000200, 1)
		return true
	} else {
		p.SendPinkNotice("ARAN_TUTORIAL_EXIT")
		return false
	}
}
