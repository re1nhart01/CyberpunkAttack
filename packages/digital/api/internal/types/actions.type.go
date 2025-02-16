package types

import "encoding/json"

type ParsedActionsCardType struct {
	CardList []ActionCardCommonType `json:"as_list"`
	Actions  struct {
		Hacks   []ActionCardCommonType `json:"hacks"`
		Boom    []ActionCardCommonType `json:"boom"`
		Heals   []ActionCardCommonType `json:"heals"`
		Runners []ActionCardCommonType `json:"runners"`
		Evades  []ActionCardCommonType `json:"evades"`
		Gun     []ActionCardCommonType `json:"gun"`
		Other   []ActionCardCommonType `json:"other"`
	}
}

//{
//      "id": 1,
//      "type": "actions",
//      "action_type": "hacks",
//      "name": "Script",
//      "spec": "action_hack_script",
//      "description": "The selected opponent discards (1) random card",
//      "illustration": "script_ill",
//      "target": "single",
//      "payload": [
//        {
//          "action": "discard",
//          "what": "deck",
//          "howMuch": 1,
//          "target": "target"
//        }
//      ]
//    },

const (
	ActionType       = "actions"
	ActionTypeHacks  = "hacks"
	ActionTypeBoom   = "boom"
	ActionTypeHeals  = "heals"
	ActionTypeEvades = "evades"
	ActionTypeGun    = "gun"
	ActionTypeOther  = "other"
)

// all specs
const (
	ActionScriptSpec = "action_hack_script"
)

type ActionCardCommonType struct {
	Id           int                     `json:"id"`
	Type         string                  `json:"type"`
	ActionType   string                  `json:"action_type"`
	Name         string                  `json:"name"`
	Spec         string                  `json:"spec"`
	Description  string                  `json:"description"`
	Illustration string                  `json:"illustration"`
	Target       string                  `json:"target"`
	Payload      []ActionCardPayloadType `json:"payload"`
}

type ActionCardPayloadType struct {
}

func (t *ParsedActionsCardType) fromJson(d []byte) *ParsedActionsCardType {
	data := ParsedActionsCardType{}
	err := json.Unmarshal(d, &data)
	if err != nil {
		return &ParsedActionsCardType{}
	}
	return &data
}
