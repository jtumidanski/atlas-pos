package portal

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func getId(m Model) (uint32, error) {
	return m.Id(), nil
}

func ByNamePortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, name string) model.IdProvider[uint32] {
	return func(mapId uint32, name string) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter(ByNameModelProvider(l, span)(mapId, name), getId)
	}
}

func RandomPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.IdProvider[uint32] {
	return func(mapId uint32) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter(randomPortalModelProvider(l, span)(mapId), getId)
	}
}

func randomPortalModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return model.SliceProviderToProviderAdapter(ByMapModelListProvider(l, span)(mapId), model.RandomPreciselyOneFilter[Model])
	}
}

func MarketPreciselyOneFilter(models []Model) (Model, error) {
	if len(models) == 0 {
		return Model{}, errors.New("none found")
	}
	for _, p := range models {
		if p.ScriptName() != "" && strings.Contains(p.ScriptName(), "market") {
			return p, nil
		}
	}
	return model.RandomPreciselyOneFilter(models)
}

func MarketPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.IdProvider[uint32] {
	return func(mapId uint32) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter(marketModelProvider(l, span)(mapId), getId)
	}
}

func marketModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return model.SliceProviderToProviderAdapter(ByMapModelListProvider(l, span)(mapId), MarketPreciselyOneFilter)
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalId uint32) model.Provider[Model] {
	return func(mapId uint32, portalId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestById(mapId, portalId), makeModel)
	}
}

func ByNameModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalName string) model.Provider[Model] {
	return func(mapId uint32, portalName string) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestByName(mapId, portalName), makeModel)
	}
}

func ByMapModelListProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.SliceProvider[Model] {
	return func(mapId uint32) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestAll(mapId), makeModel)
	}
}

func makeModel(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	attr := body.Attributes
	return NewPortalModel(uint32(id), attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName), nil
}

func WarpToPortal(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p model.IdProvider[uint32]) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p model.IdProvider[uint32]) {
		emitChangeMap(l, span)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpRandom(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, RandomPortalIdProvider(l, span)(mapId))
	}
}

func WarpById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, model.FixedIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, ByNamePortalIdProvider(l, span)(mapId, portalName))
	}
}
