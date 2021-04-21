package script

import (
	"github.com/sirupsen/logrus"
)

type infoMiniMap struct {
}

func InfoMiniMap() PortalScript {
	return infoMiniMap{}
}

func (a infoMiniMap) Name() string {
	return "infoMinimap"
}

func (a infoMiniMap) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.QuestStarted(1031) {
		p.ShowInfo("UI/tutorial.img/25")
	}
	p.BlockPortal()
	return true
}
