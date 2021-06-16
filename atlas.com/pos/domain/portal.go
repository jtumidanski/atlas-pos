package domain

type PortalModel struct {
	id          uint32
	name        string
	target      string
	theType     uint8
	x           int16
	y           int16
	targetMapId uint32
	scriptName  string
}

func (p PortalModel) ScriptName() string {
	return p.scriptName
}

func (p PortalModel) TargetMapId() uint32 {
	return p.targetMapId
}

func (p PortalModel) Target() string {
	return p.target
}

func (p PortalModel) Id() uint32 {
	return p.id
}

func (p PortalModel) Name() string {
	return p.name
}

func NewPortalModel(id uint32, name string, target string, targetMapId uint32, theType uint8, x int16, y int16, scriptName string) PortalModel {
	return PortalModel{
		id:          id,
		name:        name,
		target:      target,
		targetMapId: targetMapId,
		theType:     theType,
		x:           x,
		y:           y,
		scriptName:  scriptName,
	}
}
