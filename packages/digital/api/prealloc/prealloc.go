package prealloc

import (
	"encoding/json"
	"fmt"
	embed "github.com/cyberpunkattack"
	"types"
)

var ActionsMap map[string]*types.ActionCardCommonType
var ImplantsMap map[string]*types.ImplantCardCommonType

func init() {
	fmt.Println("prealloc init")
	ToJSONActions(embed.EmbedActionsMap, ActionsMap)
	ToJSONImplants(embed.EmbedImplantsMap, ImplantsMap)
}

func ToJSONActions(d []byte, to map[string]*types.ActionCardCommonType) *map[string]*types.ActionCardCommonType {
	result := &map[string]*types.ActionCardCommonType{}

	if err := json.Unmarshal(d, &to); err != nil {
		//TODO: make logger
		return result
	}

	return result
}

func ToJSONImplants(d []byte, to map[string]*types.ImplantCardCommonType) *map[string]*types.ImplantCardCommonType {
	result := &map[string]*types.ImplantCardCommonType{}

	if err := json.Unmarshal(d, &to); err != nil {
		//TODO: make logger
		return result
	}

	return result
}
