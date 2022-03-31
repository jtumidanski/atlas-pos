package portal

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
)

const DefaultPortalId uint32 = 0

type IdProvider func() uint32

type PreciselyOneFilter func([]Model) (Model, error)

func FixedPortalIdProvider(portalId uint32) IdProvider {
	return func() uint32 {
		return portalId
	}
}

func DefaultIdProvider() uint32 {
	return DefaultPortalId
}

func FromPortalIdProvider(p Model) IdProvider {
	return func() uint32 {
		return p.Id()
	}
}

func modelProviderToIdProviderAdapter(l logrus.FieldLogger) func(mp model.Provider[Model]) IdProvider {
	return func(mp model.Provider[Model]) IdProvider {
		p, err := mp()
		if err != nil {
			l.WithError(err).Warnf("Unable to retrieve portal. Using default.")
			return DefaultIdProvider
		}
		return FromPortalIdProvider(p)
	}
}

func ByNamePortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, name string) IdProvider {
	return func(mapId uint32, name string) IdProvider {
		return modelProviderToIdProviderAdapter(l)(ByNameModelProvider(l, span)(mapId, name))
	}
}

func modelListProviderToModelProviderAdapter(provider model.SliceProvider[Model], preciselyOneFilter PreciselyOneFilter) model.Provider[Model] {
	return func() (Model, error) {
		ps, err := provider()
		if err != nil {
			return Model{}, err
		}
		return preciselyOneFilter(ps)
	}
}

func RandomPreciselyOneFilter(models []Model) (Model, error) {
	if len(models) == 0 {
		return Model{}, errors.New("none found")
	}
	return models[rand.Intn(len(models))], nil
}

func RandomPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) IdProvider {
	return func(mapId uint32) IdProvider {
		return modelProviderToIdProviderAdapter(l)(randomPortalModelProvider(l, span)(mapId))
	}
}

func randomPortalModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return modelListProviderToModelProviderAdapter(ByMapModelListProvider(l, span)(mapId), RandomPreciselyOneFilter)
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
	return RandomPreciselyOneFilter(models)
}

func MarketPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) IdProvider {
	return func(mapId uint32) IdProvider {
		return modelProviderToIdProviderAdapter(l)(marketModelProvider(l, span)(mapId))
	}
}

func marketModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return modelListProviderToModelProviderAdapter(ByMapModelListProvider(l, span)(mapId), MarketPreciselyOneFilter)
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

func WarpToPortal(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p IdProvider) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p IdProvider) {
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
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, ByNamePortalIdProvider(l, span)(mapId, portalName))
	}
}
