package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Davy3Hd1 struct {
}

func (p Davy3Hd1) Name() string {
	return "davy3_hd1"
}

func (p Davy3Hd1) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//int level = eim.getProperty("level").toInteger()
	//if(eim.getProperty("stage3b") == "0") {
	//	pi.getMap(925100302).spawnAllMonstersFromMapSpawnList(level, true)
	//	eim.setProperty("stage3b", "1")
	//}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(925100302, 0)
	return true
}
