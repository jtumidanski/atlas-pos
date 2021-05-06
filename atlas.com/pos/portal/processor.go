package portal

import (
	"atlas-pos/domain"
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Processor struct {
	l logrus.FieldLogger
}

func NewProcessor(l logrus.FieldLogger) *Processor {
	return &Processor{l}
}

func (p *Processor) GetMapPortalById(mapId uint32, portalId uint32) (*domain.PortalModel, error) {
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

func (p *Processor) GetMapPortalByName(mapId uint32, portalName string) (*domain.PortalModel, error) {
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

func makePortal(id uint32, attr attributes.PortalAttributes) domain.PortalModel {
	return domain.NewPortalModel(id, attr.Name, attr.Target, attr.TargetMap, attr.Type, attr.X, attr.Y, attr.ScriptName)
}
