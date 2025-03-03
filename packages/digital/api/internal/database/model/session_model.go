package models

import (
	"time"
	"types"
)

type ROLES_UNION string
type ActionDeckType []*types.ActionCardCommonType
type ImplantDeckType []*types.ImplantCardCommonType
type RolesDeckType []*types.RoleCardCommonType

const (
	CYBERPUNK_ROLE      ROLES_UNION = "cyberpunk"
	CORPORATE_ROLE      ROLES_UNION = "corporate"
	CYBERSTRAY_ROLE     ROLES_UNION = "cyberstray"
	CYBERPUNK_BOSS_ROLE ROLES_UNION = "cyberpunk_boss"
	CORPORATE_BOSS_ROLE ROLES_UNION = "corporate_boss"
)

const (
	PENDING_STATUS = "pending"
	ACTIVE_STATUS  = "active"
	ENDED_STATUS   = "ended"
	FALSE_STATUS   = "false"
	ERROR_STATUS   = "error"
)

const (
	//GAME_TYPE_RANKED = "ranked"
	GAME_TYPE_RANDOM = "random"
	GAME_TYPE_CUSTOM = "custom"
)

type SessionMoveModel struct {
}

type SessionPG struct {
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	SessionID       string       `json:"session_id"`
	UserIds         []string     `json:"userIds"`
	Winner          string       `json:"winner"`
	CreatorHash     string       `json:"creator_hash"`
	WinnerRole      ROLES_UNION  `json:"winnerRole"`
	WithError       string       `json:"with_error"`
	EndedGracefully bool         `json:"ended_gracefully"`
	Flags           SessionFlags `json:"flags"`
	CreateAt        time.Time    `json:"createAt"`
	EndedAt         time.Time    `json:"endedAt"`
}

type SessionFlags struct {
	Started bool `json:"started"`
}

type SessionCommitment struct {
	ActionDeck  []ActionDeckType            `json:"action_deck"`
	MovesList   []SessionMoveModel          `json:"session_move_list"`
	ImplantDeck []ImplantDeckType           `json:"implant_deck"`
	Flags       SessionFlags                `json:"flags"`
	UserMap     map[string]UserSessionModel `json:"user_map"`
	Password    string                      `json:"password"`
}

type SessionIM struct {
	ID          int                         `json:"id"`
	Name        string                      `json:"name"`
	CreatorHash string                      `json:"creator_hash"`
	SessionID   string                      `json:"session_id"`
	Password    string                      `json:"password"`
	UserIds     []string                    `json:"user_ids"`
	UserMap     map[string]UserSessionModel `json:"user_map"`
	ActionDeck  []ActionDeckType            `json:"action_deck"`
	RoleDeck    []RolesDeckType             `json:"role_deck"`
	MovesList   []SessionMoveModel          `json:"session_move_list"`
	ImplantDeck []ImplantDeckType           `json:"implant_deck"`
	SideEffects []SideEffect                `json:"side_effects"`
	Type        string                      `json:"type"`
	Status      string                      `json:"status"`
	IsEnded     bool                        `json:"is_ended"`
	CreatedAt   time.Time                   `json:"created_at"`
	EndedAt     time.Time                   `json:"ended_at"`
	MovesCount  int16                       `json:"moves_count"`
	Flags       SessionFlags                `json:"flags"`
}

type SideEffect struct {
}

type UserEffectsModel struct {
	IsHealDisabled     bool `json:"is_heal_disabled"`
	IsImplantsDisabled bool `json:"is_implants_disabled"`
}

type UserSessionModel struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	UserHash    string            `json:"userHash"`
	Health      uint8             `json:"health"`
	Role        ROLES_UNION       `json:"role"`
	Implants    []ImplantDeckType `json:"implants"`
	Actions     []ActionDeckType  `json:"actions"`
	UserEffects UserEffectsModel  `json:"userEffects"`
	Dead        bool              `json:"dead"`
	CreatedAt   time.Time         `json:"createdAt"`
}
