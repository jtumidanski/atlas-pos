package portal

import (
	"atlas-pos/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	mapInformationServicePrefix string = "/ms/mis/"
	mapInformationService              = requests.BaseRequest + mapInformationServicePrefix
	mapsResource                       = mapInformationService + "maps/"
	portalsResource                    = mapsResource + "%d/portals"
	portalsByName                      = portalsResource + "?name=%s"
	portalResource                     = portalsResource + "/%d"
)

type Request func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l, span)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestById(mapId uint32, portalId uint32) Request {
	return makeRequest(fmt.Sprintf(portalResource, mapId, portalId))
}

func requestByName(mapId uint32, portalName string) Request {
	return makeRequest(fmt.Sprintf(portalsByName, mapId, portalName))
}

func requestAll(mapId uint32) Request {
	return makeRequest(fmt.Sprintf(portalsResource, mapId))
}
