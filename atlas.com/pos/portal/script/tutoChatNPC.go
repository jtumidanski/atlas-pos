package script

import (
	"github.com/sirupsen/logrus"
)

type tutoChatNPC struct {
}

func TutoChatNPC() PortalScript {
	return tutoChatNPC{}
}

func (a tutoChatNPC) Name() string {
	return "tutoChatNPC"
}

func (a tutoChatNPC) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.HasLevel30Character() {
		p.OpenNPC(2007)
	}
	p.BlockPortal()
	return true
}
