package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Davy2Hd1 struct {
}

func (p Davy2Hd1) Name() string {
	return "davy2_hd1"
}

func (p Davy2Hd1) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//int level = eim.getProperty("level").toInteger()
	//if (eim.getProperty("stage2b") == "0") {
	//	pi.getMap(925100202).spawnAllMonstersFromMapSpawnList(level, true)
	//	eim.setProperty("stage2b", "1")
	//}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(925100202, 0)
	return true
}
