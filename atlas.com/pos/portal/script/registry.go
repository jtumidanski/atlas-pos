package script

import (
	"errors"
	"sync"
)

type ScriptRegistry struct {
	registry map[string]Script
}

var once sync.Once
var scriptRegistry *ScriptRegistry

func GetScriptRegistry() *ScriptRegistry {
	once.Do(func() {
		scriptRegistry = initRegistry()
	})
	return scriptRegistry
}

func initRegistry() *ScriptRegistry {
	s := &ScriptRegistry{make(map[string]Script)}
	s.addPortalScript(ChristmasOut2008())
	s.addPortalScript(Advice00())
	s.addPortalScript(Advice01())
	s.addPortalScript(Advice02())
	s.addPortalScript(Advice03())
	s.addPortalScript(Advice04())
	s.addPortalScript(Advice05())
	s.addPortalScript(Advice06())
	s.addPortalScript(Advice07())
	s.addPortalScript(Advice08())
	s.addPortalScript(Advice09())
	s.addPortalScript(AdviceMap())
	s.addPortalScript(AMatchMove2())
	s.addPortalScript(Apq00())
	s.addPortalScript(Apq01())
	s.addPortalScript(Apq1())
	s.addPortalScript(Apq02())
	s.addPortalScript(Apq2())
	s.addPortalScript(Apq3())
	s.addPortalScript(ApqClosed())
	s.addPortalScript(ApqDoor())
	s.addPortalScript(AquaPQBoss0())
	s.addPortalScript(AranTutorAloneX())
	s.addPortalScript(AranTutorArrow0())
	s.addPortalScript(AranTutorArrow1())
	s.addPortalScript(AranTutorArrow2())
	s.addPortalScript(AranTutorArrow3())
	s.addPortalScript(AranTutorGuide0())
	s.addPortalScript(AranTutorGuide1())
	s.addPortalScript(AranTutorGuide2())
	s.addPortalScript(AranTutorLost())
	s.addPortalScript(AranTutorMono0())
	s.addPortalScript(AranTutorMono1())
	s.addPortalScript(AranTutorMono2())
	s.addPortalScript(AranTutorMono3())
	s.addPortalScript(AranTutorOut1())
	s.addPortalScript(AranTutorOut2())
	s.addPortalScript(AranTutorOut3())
	s.addPortalScript(AriantAgit())
	s.addPortalScript(AriantCastle())
	s.addPortalScript(AriantQueens())
	s.addPortalScript(AriantMount())
	s.addPortalScript(AriantMount2())
	s.addPortalScript(BabyPigOut())
	s.addPortalScript(BalogEnd())
	s.addPortalScript(BalogTemple())
	s.addPortalScript(BedroomOut())
	s.addPortalScript(CannonTuto06())
	s.addPortalScript(CannonTuto07())
	s.addPortalScript(CannonTuto09())
	s.addPortalScript(CannonTuto10())
	s.addPortalScript(CaptinsG00())
	s.addPortalScript(CatPriestMap())
	s.addPortalScript(ContactDragon())
	s.addPortalScript(CurseForest())
	s.addPortalScript(DavyNext0())
	s.addPortalScript(DavyNext1())
	s.addPortalScript(DavyNext2())
	s.addPortalScript(EnterAchter())
	s.addPortalScript(EnterMagiclibrar())
	s.addPortalScript(TutoChatNPC())
	s.addPortalScript(InfoMiniMap())
	s.addPortalScript(GlTutoMsg0())
	return s
}

func (s *ScriptRegistry) GetScript(name string) (*Script, error) {
	if val, ok := s.registry[name]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *ScriptRegistry) addPortalScript(handler PortalScript) {
	s.registry[handler.Name()] = NewScript(handler)
}
