package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DavyNext1 struct {
}

func (p DavyNext1) Name() string {
	return "davy_next1"
}

func (p DavyNext1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.GetEventProperty(l, c)("stage2") == "3" {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(925100200, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
