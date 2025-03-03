package types

type CardsCounterType struct {
	Actions  map[string]int `json:"actions"`
	Implants map[string]int `json:"implants"`
}

func (deck *CardsCounterType) GetByDeck(deckName string) map[string]int {
	switch deckName {
	case "actions":
		return deck.Actions
	case "implants":
		return deck.Actions
	default:
		return map[string]int{}
	}
}

type RolesCounterType struct {
	Two   map[string]int `json:"2"`
	Three map[string]int `json:"3"`
	Four  map[string]int `json:"4"`
	Five  map[string]int `json:"5"`
	Six   map[string]int `json:"6"`
	Seven map[string]int `json:"7"`
	Eight map[string]int `json:"8"`
}

func (roles *RolesCounterType) GetByCount(count int) map[string]int {
	switch count {
	case 2:
		return roles.Two
	case 3:
		return roles.Three
	case 4:
		return roles.Four
	case 5:
		return roles.Five
	case 6:
		return roles.Six
	case 7:
		return roles.Seven
	case 8:
		return roles.Eight
	default:
		return map[string]int{}
	}
}
