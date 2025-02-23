package gameplay

import (
	"encoding/json"
	"github.com/cyberpunkattack/pkg/wstorify"
)

type EVENTS string
type DECK string

const (
	ActionDeck DECK = "actions"
	Implants   DECK = "implants"
)

type GameEssentialEvent wstorify.EssentialEvent

const (
	GameStarted           EVENTS = "game_started"
	SelectedRole          EVENTS = "selected_role"
	RoleSelect            EVENTS = "select_role"
	RoleSelectionComplete EVENTS = "role_selection_complete"
	StartTurn             EVENTS = "start_turn"
	ActonPerformed        EVENTS = "action_performed"
	Surrend               EVENTS = "surrend_game"
	UpdateGameState       EVENTS = "update_game_state"
	ImportantGameEvent    EVENTS = "game_event_important"
	InconsistentGame      EVENTS = "inconsisted_game"
	PauseStart            EVENTS = "start_pause"
	GamePauseEnd          EVENTS = "end_pause"
	RoundEnded            EVENTS = "round_ended"
	RoundStarted          EVENTS = "round_started"
	GameEnd               EVENTS = "game_end"
	UserReconnect         EVENTS = "reconnect_user"
	GamePaused            EVENTS = "game_paused"
	UserDisconnect        EVENTS = "disconnected_user"
	UserReconnected       EVENTS = "reconnected_user"
)

func (essEv *GameEssentialEvent) toJSON() string {
	v, err := json.Marshal(&essEv)
	if err != nil {
		return err.Error()
	}

	return string(v)
}

func (event EVENTS) CreateGameStartedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateSelectedRoleEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateRoleSelectEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateRoleSelectionCompleteEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateStartTurnEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateActionPerformedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateSurrenderEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateUpdateGameStateEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateImportantGameEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateInconsistentGameEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreatePauseStartEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateGamePauseEndEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateRoundEndedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateRoundStartedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateGameEndEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateUserReconnectEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateGamePausedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateUserDisconnectEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}

func (event EVENTS) CreateUserReconnectedEvent() *GameEssentialEvent {
	return &GameEssentialEvent{}
}
