package prealloc

import (
	"encoding/json"
	"fmt"
	"types"

	embed "github.com/cyberpunkattack"
)

var ActionsMap map[string]*types.ActionCardCommonType
var ImplantsMap map[string]*types.ImplantCardCommonType
var RolesMap map[string]*types.RoleCardCommonType
var CardsCounter *types.CardsCounterType
var RolesCounter *types.RolesCounterType

func init() {
	fmt.Println("prealloc init")
	ToJSONActions(embed.EmbedActionsMap, ActionsMap)
	ToJSONImplants(embed.EmbedImplantsMap, ImplantsMap)
	ToJSONCardsCounter(embed.EmbedActionsCounter, embed.EmbedImplantsCounter, CardsCounter)
	ToJSONRolesCounter(embed.EmbedRolesCounter, RolesCounter)
	ToJsonRolesMap(embed.EmbedRolesMap, RolesMap)
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

func ToJSONCardsCounter(d1 []byte, d2 []byte, to *types.CardsCounterType) *types.CardsCounterType {
	result := &types.CardsCounterType{}

	if err := json.Unmarshal(d1, &to.Actions); err != nil {
		//TODO: make logger
		return result
	}

	if err := json.Unmarshal(d2, &to.Actions); err != nil {
		//TODO: make logger
		return result
	}

	return result
}

func ToJSONRolesCounter(d []byte, to *types.RolesCounterType) *types.RolesCounterType {
	result := &types.RolesCounterType{}

	if err := json.Unmarshal(d, to); err != nil {
		//TODO: make logger
		return result
	}

	return result
}

func ToJsonRolesMap(d []byte, to map[string]*types.RoleCardCommonType) map[string]*types.RoleCardCommonType {
	result := map[string]*types.RoleCardCommonType{}

	if err := json.Unmarshal(d, &to); err != nil {
		//TODO: make logger
		return result
	}

	return result
}
