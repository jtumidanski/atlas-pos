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

func requestById(mapId uint32, portalId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(portalResource, mapId, portalId))
}

func requestByName(mapId uint32, portalName string) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(portalsByName, mapId, portalName))
}

func requestAll(mapId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(portalsResource, mapId))
}
