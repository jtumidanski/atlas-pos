package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3RoomOut struct {
}

func (p Party3RoomOut) Name() string {
	return "party3_roomout"
}

func (p Party3RoomOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	exitPortal := uint32(0)

	switch c.MapId() {
	case 920010200:
		exitPortal = 4
		break
	case 920010300:
		exitPortal = 12
		break
	case 920010400:
		exitPortal = 5
		break
	case 920010500:
		exitPortal = 13
		break
	case 920010600:
		exitPortal = 15
		break
	case 920010700:
		exitPortal = 14
		break
	case 920011000:
		exitPortal = 16
		break
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010100, exitPortal)
	return true
}
