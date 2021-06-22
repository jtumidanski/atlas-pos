package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Zakum03 struct {
}

func (p Zakum03) Name() string {
	return "Zakum03"
}

func (p Zakum03) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if (!pi.getEventInstance().isEventCleared()) {
	//	pi.sendPinkNotice("ZAKUM_COMPLETE_TRIALS")
	//	return false
	//}
	//
	//if (pi.getEventInstance().gridCheck(pi.getCharacterId()) == -1) {
	//	pi.sendPinkNotice("ZAKUM_CLAIM_PRIZE")
	//	return false
	//}

	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(211042300)
	return true
}
