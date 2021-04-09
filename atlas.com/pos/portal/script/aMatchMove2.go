package script

import (
	"log"
)

type aMatchMove2 struct {
}

func AMatchMove2() PortalScript {
	return aMatchMove2{}
}

func (a aMatchMove2) Name() string {
	return "aMatchMove2"
}

func (a aMatchMove2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(p.GetSavedLocation("MIRROR"))
	return true
}
