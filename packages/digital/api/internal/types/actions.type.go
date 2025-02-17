package types

import (
	"encoding/json"
	embed "github.com/cyberpunkattack"
)

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

//"buff": {
//"name": "two_or_more_implants_two_damage",
//"action": "attack",
//"howMuch": 2,
//"reason": "two_or_more_implants"
//},

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
	Id           int                     `json:"id"`           //id of card (int)
	Type         string                  `json:"type"`         //actions | implant | virus
	ActionType   string                  `json:"action_type"`  // hacks | ...
	Name         string                  `json:"name"`         // name of card
	Spec         string                  `json:"spec"`         // special identifier
	Description  string                  `json:"description"`  // description of card
	Illustration string                  `json:"illustration"` // illustration id
	Target       string                  `json:"target"`       // simple | mass
	Payload      []ActionCardPayloadType `json:"payload"`
}

type ActionCardPayloadType struct {
	Action       string         `json:"action"`         // action of payload, like discard, boom and etc
	What         string         `json:"what"`           // deck and etc
	HowMuch      string         `json:"howMuch"`        // count of card, damage etc.
	Target       string         `json:"target"`         // to whom to shot or discard
	Health       string         `json:"health"`         // health of runner
	Buff         map[string]any `json:"buff"`           // buff for some card
	Debuff       map[string]any `json:"debuff"`         //debuff for some card
	Range        string         `json:"range"`          // range of action
	DealDamage   string         `json:"deal_damage"`    // object with shots: 1, damage: 2
	EvadeType    string         `json:"evade_type"`     //  "evade_type": ["boom", "gun", "combat_implant"],
	IsReturnBack bool           `json:"is_return_back"` // is shot or something else returns back
	CanBeDodget  bool           `json:"can_be_dodget"`  // can be dodged (for shots)
	Deck         string         `json:"deck"`           // for example when draw, then from which deck draw
	BackOnly     bool           `json:"back_only"`      // return_damage_back back_only: ["gun"]
}

func (t *ParsedActionsCardType) fromJson(d []byte) *ParsedActionsCardType {
	data := ParsedActionsCardType{}
	err := json.Unmarshal(d, &data)
	if err != nil {
		return &ParsedActionsCardType{}
	}
	return &data
}

func a() {
	data := &ParsedActionsCardType{}
	data.fromJson(embed.EmbedActionsDeck)
}
