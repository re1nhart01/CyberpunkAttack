package gameplay

import (
	"context"
	models "github.com/cyberpunkattack/database/model"
	"time"
)

// selector - file where you will select -> actions | cyberimplants -> which category -> run specific handler for every card category

type ActionObject struct {
}

type ExecutionObject struct {
	Creator   string           `json:"creator"`
	SessionId string           `json:"sessionId"`
	Session   models.SessionIM `json:"session"`
	Now       time.Time        `json:"now"`
	Action    ActionObject     `json:"action"`
}

func HandleSelectDeck(ctx context.Context, executionObj *ExecutionObject) {

}

func HandleSelectCategory(ctx context.Context, executionObj *ExecutionObject) {

}
