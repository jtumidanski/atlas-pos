package portal

import (
	"atlas-pos/rest/requests"
	"fmt"
)

const (
	mapInformationServicePrefix string = "/ms/mis/"
	mapInformationService              = requests.BaseRequest + mapInformationServicePrefix
	mapsResource                       = mapInformationService + "maps/"
	portalsResource                    = mapsResource + "%d/portals"
	portalsByName                      = portalsResource + "?name=%s"
	portalResource                     = portalsResource + "/%d"
)

func requestById(mapId uint32, portalId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(portalResource, mapId, portalId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestByName(mapId uint32, portalName string) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(portalsByName, mapId, portalName), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestAll(mapId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(portalsResource, mapId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
