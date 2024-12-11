package models

import "time"


type ROLES_UNION string
type ActionDeckType map[string]any
type ImplantDeckType map[string]any


const  (
	CYBERPUNK_ROLE ROLES_UNION = "cyberpunk"
	CORPORATE_ROLE ROLES_UNION = "corporate"
	CORPORATE_BOSS_ROLE ROLES_UNION = "corporate_boss"
	CYBERSTRAY_ROLE ROLES_UNION = "cyberstray"
)

const (
	GAME_TYPE_RANKED = "ranked"
	GAME_TYPE_RANDOM = "random"
	GAME_TYPE_CUSTOM = "custom"
)

type SessionPG struct {
	ID 			int 	  		`json:"id"`
	Name 		string 	  		`json:"name"`
	UserIds 	[]string 	    `json:"userIds"`
	Winner 		string 			`json:"winner"`
	WinnerRole  ROLES_UNION     `json:"winnerRole"`
	CreateAt 	time.Time 	  	`json:"createAt"`
	EndedAt 	time.Time 	  	`json:"endedAt"`
}

type SessionIM struct {
	ID 			int 	  					 `json:"id"`
	Name 		string 	  					 `json:"name"`
	Password 	string 						 `json:"password"`
	UserIds 	[]string 	  				 `json:"userIds"`
	UserMap 	map[string]UserSessionModel  `json:"userMap"`
	ActionDeck 	[]ActionDeckType			 `json:"actionDeck"`
	ImplantDeck []ImplantDeckType 		   	 `json:"implantDeck"`
	Type        string                       `json:"type"`
	IsEnded     bool                         `json:"isEnded"`
	CreateAt 	time.Time 	  				 `json:"createAt"`
	EndedAt 	time.Time 	  				 `json:"endedAt"`
}

type UserSessionModel struct {
	Id 			int 				`json:"id"`
	Name 		string 				`json:"name"`
	UserHash 	string 				`json:"userHash"`
	Health 		uint8 				`json:"health"`
	Role 		ROLES_UNION 		`json:"role"`
	Implants 	[]ImplantDeckType 	`json:"implants"`
	Actions 	[]ActionDeckType 	`json:"actions"`
	Dead 		bool 				`json:"dead"`
	CreatedAt 	time.Time 			`json:"createdAt"`
}

