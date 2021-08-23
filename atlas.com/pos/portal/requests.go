package portal

import (
	"atlas-pos/rest/requests"
	"fmt"
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

func request(l logrus.FieldLogger) func(url string) (*dataContainer, error) {
	return func(url string) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestById(l logrus.FieldLogger) func(mapId uint32, portalId uint32) (*dataContainer, error) {
	return func(mapId uint32, portalId uint32) (*dataContainer, error) {
		return request(l)(fmt.Sprintf(portalResource, mapId, portalId))
	}
}

func requestByName(l logrus.FieldLogger) func(mapId uint32, portalName string) (*dataContainer, error) {
	return func(mapId uint32, portalName string) (*dataContainer, error) {
		return request(l)(fmt.Sprintf(portalsByName, mapId, portalName))
	}
}

func requestAll(l logrus.FieldLogger) func(mapId uint32) (*dataContainer, error) {
	return func(mapId uint32) (*dataContainer, error) {
		return request(l)(fmt.Sprintf(portalsResource, mapId))
	}
}
