package script

import (
	"log"
)

type ariantAgit struct {
}

func AriantAgit() PortalScript {
	return ariantAgit{}
}

func (a ariantAgit) Name() string {
	return "ariant_Agit"
}

func (a ariantAgit) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	if p.QuestCompleted(3928) && p.QuestCompleted(3931) && p.QuestCompleted(3934) {
		p.PlayPortalSound()
		p.WarpById(260000201, 1)
		return true
	} else {
		p.SendPinkNotice("SAND_BANDITS_ONLY")
		return false
	}
}
