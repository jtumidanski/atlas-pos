package registry

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/discrete"
	"errors"
	"sync"
)

type ScriptRegistry struct {
	registry map[string]script.Script
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
	s := &ScriptRegistry{make(map[string]script.Script)}
	s.addPortalScript(discrete.ChristmasOut2008{})
	s.addPortalScript(discrete.Advice00{})
	s.addPortalScript(discrete.Advice01{})
	s.addPortalScript(discrete.Advice02{})
	s.addPortalScript(discrete.Advice03{})
	s.addPortalScript(discrete.Advice04{})
	s.addPortalScript(discrete.Advice05{})
	s.addPortalScript(discrete.Advice06{})
	s.addPortalScript(discrete.Advice07{})
	s.addPortalScript(discrete.Advice08{})
	s.addPortalScript(discrete.Advice09{})
	s.addPortalScript(discrete.AdviceMap{})
	s.addPortalScript(discrete.AMatchMove2{})
	s.addPortalScript(discrete.Apq00{})
	s.addPortalScript(discrete.Apq01{})
	s.addPortalScript(discrete.Apq1{})
	s.addPortalScript(discrete.Apq02{})
	s.addPortalScript(discrete.Apq2{})
	s.addPortalScript(discrete.Apq3{})
	s.addPortalScript(discrete.ApqClosed{})
	s.addPortalScript(discrete.ApqDoor{})
	s.addPortalScript(discrete.AquaPQBoss0{})
	s.addPortalScript(discrete.AranTutorAloneX{})
	s.addPortalScript(discrete.AranTutorArrow0{})
	s.addPortalScript(discrete.AranTutorArrow1{})
	s.addPortalScript(discrete.AranTutorArrow2{})
	s.addPortalScript(discrete.AranTutorArrow3{})
	s.addPortalScript(discrete.AranTutorGuide0{})
	s.addPortalScript(discrete.AranTutorGuide1{})
	s.addPortalScript(discrete.AranTutorGuide2{})
	s.addPortalScript(discrete.AranTutorLost{})
	s.addPortalScript(discrete.AranTutorMono0{})
	s.addPortalScript(discrete.AranTutorMono1{})
	s.addPortalScript(discrete.AranTutorMono2{})
	s.addPortalScript(discrete.AranTutorMono3{})
	s.addPortalScript(discrete.AranTutorOut1{})
	s.addPortalScript(discrete.AranTutorOut2{})
	s.addPortalScript(discrete.AranTutorOut3{})
	s.addPortalScript(discrete.AriantAgit{})
	s.addPortalScript(discrete.AriantCastle{})
	s.addPortalScript(discrete.AriantQueens{})
	s.addPortalScript(discrete.AriantMount{})
	s.addPortalScript(discrete.AriantMount2{})
	s.addPortalScript(discrete.BabyPigOut{})
	s.addPortalScript(discrete.BalogEnd{})
	s.addPortalScript(discrete.BalogTemple{})
	s.addPortalScript(discrete.BedroomOut{})
	s.addPortalScript(discrete.CannonTuto06{})
	s.addPortalScript(discrete.CannonTuto07{})
	s.addPortalScript(discrete.CannonTuto09{})
	s.addPortalScript(discrete.CannonTuto10{})
	s.addPortalScript(discrete.CaptinsG00{})
	s.addPortalScript(discrete.CatPriestMap{})
	s.addPortalScript(discrete.ContactDragon{})
	s.addPortalScript(discrete.CurseForest{})
	s.addPortalScript(discrete.Davy2Hd1{})
	s.addPortalScript(discrete.Davy3Hd1{})
	s.addPortalScript(discrete.DavyNext0{})
	s.addPortalScript(discrete.DavyNext1{})
	s.addPortalScript(discrete.DavyNext2{})
	s.addPortalScript(discrete.DavyNext3{})
	s.addPortalScript(discrete.DavyNext4{})
	s.addPortalScript(discrete.DepartGoBack00{})
	s.addPortalScript(discrete.DepartGoBack01{})
	s.addPortalScript(discrete.DepartGoFoward0{})
	s.addPortalScript(discrete.DepartGoFoward1{})
	s.addPortalScript(discrete.DepartToKerning{})
	s.addPortalScript(discrete.DepartTopFloor{})
	s.addPortalScript(discrete.DepartTopOut{})
	s.addPortalScript(discrete.DojangExit{})
	s.addPortalScript(discrete.DojangNext{})
	s.addPortalScript(discrete.DojangTuto{})
	s.addPortalScript(discrete.DojangUp{})
	s.addPortalScript(discrete.Dracoout{})
	s.addPortalScript(discrete.DragonEggNotice{})
	s.addPortalScript(discrete.DragonEyes{})
	s.addPortalScript(discrete.DragonNest{})
	s.addPortalScript(discrete.EBoat1{})
	s.addPortalScript(discrete.EBoat2{})
	s.addPortalScript(discrete.Elevator{})
	s.addPortalScript(discrete.ElizaGarden{})
	s.addPortalScript(discrete.EndCow{})
	s.addPortalScript(discrete.EnterEarth00{})
	s.addPortalScript(discrete.EnterEarth01{})
	s.addPortalScript(discrete.EnterNautil{})
	s.addPortalScript(discrete.EnterTd{})
	s.addPortalScript(discrete.EnterAchter{})
	s.addPortalScript(discrete.EnterBackStreet{})
	s.addPortalScript(discrete.EnterDisguise0{})
	s.addPortalScript(discrete.EnterDisguise1{})
	s.addPortalScript(discrete.EnterDisguise2{})
	s.addPortalScript(discrete.EnterDisguise3{})
	s.addPortalScript(discrete.EnterDisguise4{})
	s.addPortalScript(discrete.EnterDisguise5{})
	s.addPortalScript(discrete.EnterDollcave{})
	s.addPortalScript(discrete.EnterMagiclibrar{})
	s.addPortalScript(discrete.TutoChatNPC{})
	s.addPortalScript(discrete.InfoMiniMap{})
	s.addPortalScript(discrete.GlTutoMsg0{})
	return s
}

func (s *ScriptRegistry) GetScript(name string) (*script.Script, error) {
	if val, ok := s.registry[name]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *ScriptRegistry) addPortalScript(handler script.PortalScript) {
	s.registry[handler.Name()] = script.NewScript(handler)
}
