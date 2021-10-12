package portal

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
)

const DefaultPortalId uint32 = 0

type IdProvider func() uint32

type ModelProvider func() (*Model, error)

type ModelListProvider func() ([]*Model, error)

type PreciselyOneFilter func([]*Model) (*Model, error)

func FixedPortalIdProvider(portalId uint32) IdProvider {
	return func() uint32 {
		return portalId
	}
}

func DefaultIdProvider() uint32 {
	return DefaultPortalId
}

func FromPortalIdProvider(p *Model) IdProvider {
	return func() uint32 {
		return p.Id()
	}
}

func modelProviderToIdProviderAdapter(l logrus.FieldLogger) func(mp ModelProvider) IdProvider {
	return func(mp ModelProvider) IdProvider {
		p, err := mp()
		if p == nil || err != nil {
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

func modelListProviderToModelProviderAdapter(provider ModelListProvider, preciselyOneFilter PreciselyOneFilter) ModelProvider {
	return func() (*Model, error) {
		ps, err := provider()
		if err != nil {
			return nil, err
		}
		return preciselyOneFilter(ps)
	}
}

func RandomPreciselyOneFilter(models []*Model) (*Model, error) {
	if len(models) == 0 {
		return nil, nil
	}
	return models[rand.Intn(len(models))], nil
}

func RandomPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) IdProvider {
	return func(mapId uint32) IdProvider {
		return modelProviderToIdProviderAdapter(l)(randomPortalModelProvider(l, span)(mapId))
	}
}

func randomPortalModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) ModelProvider {
	return func(mapId uint32) ModelProvider {
		return modelListProviderToModelProviderAdapter(ByMapModelListProvider(l, span)(mapId), RandomPreciselyOneFilter)
	}
}

func MarketPreciselyOneFilter(models []*Model) (*Model, error) {
	if len(models) == 0 {
		return nil, nil
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

func marketModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) ModelProvider {
	return func(mapId uint32) ModelProvider {
		return modelListProviderToModelProviderAdapter(ByMapModelListProvider(l, span)(mapId), MarketPreciselyOneFilter)
	}
}

func requestModelProvider(l logrus.FieldLogger, span opentracing.Span) func(r Request) ModelProvider {
	return func(r Request) ModelProvider {
		return func() (*Model, error) {
			resp, err := r(l, span)
			if err != nil {
				return nil, err
			}

			p, err := makeModel(resp.Data())
			if err != nil {
				return nil, err
			}
			return p, nil
		}
	}
}

func requestModelListProvider(l logrus.FieldLogger, span opentracing.Span) func(r Request) ModelListProvider {
	return func(r Request) ModelListProvider {
		return func() ([]*Model, error) {
			resp, err := r(l, span)
			if err != nil {
				return nil, err
			}

			results := make([]*Model, 0)
			for _, d := range resp.DataList() {
				p, err := makeModel(d)
				if err != nil {
					return nil, err
				}
				results = append(results, p)
			}
			return results, nil
		}
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalId uint32) ModelProvider {
	return func(mapId uint32, portalId uint32) ModelProvider {
		return requestModelProvider(l, span)(requestById(mapId, portalId))
	}
}

func ByNameModelProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalName string) ModelProvider {
	return func(mapId uint32, portalName string) ModelProvider {
		return requestModelProvider(l, span)(requestByName(mapId, portalName))
	}
}

func ByMapModelListProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) ModelListProvider {
	return func(mapId uint32) ModelListProvider {
		return requestModelListProvider(l, span)(requestAll(mapId))
	}
}

func makeModel(body *dataBody) (*Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	attr := body.Attributes
	return NewPortalModel(uint32(id), attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName), nil
}
