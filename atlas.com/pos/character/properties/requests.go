package properties

import (
	"atlas-pos/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters"
	charactersById                 = charactersResource + "/%d"
	accountCharacters              = charactersService + "?accountId=%d&worldId=%d"
)

func requestAttributesById(characterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(charactersById, characterId))
}

func requestAccountCharacters(accountId uint32, worldId byte) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(accountCharacters, accountId, worldId))
}
