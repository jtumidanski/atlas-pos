package portal

import (
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"strconv"
)

func GetMapPortalById(mapId uint32, portalId uint32) (*Model, error) {
	resp, err := requests.MapInformation().GetPortalById(mapId, portalId)
	if err != nil {
		return nil, err
	}

	d := resp.Data()
	aid, err := strconv.ParseUint(d.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	a := makePortal(uint32(aid), d.Attributes)
	return &a, nil
}

func GetMapPortalByName(mapId uint32, portalName string) (*Model, error) {
	resp, err := requests.MapInformation().GetPortalByName(mapId, portalName)
	if err != nil {
		return nil, err
	}

	d := resp.Data()
	aid, err := strconv.ParseUint(d.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	a := makePortal(uint32(aid), d.Attributes)
	return &a, nil
}

func makePortal(id uint32, attr attributes.PortalAttributes) Model {
	return NewPortalModel(id, attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName)
}
