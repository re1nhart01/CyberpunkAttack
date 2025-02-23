package gameplay

import (
	"context"
	"github.com/cyberpunkattack/pkg/wstorify"
	"types"
)

type Handler struct {
	Executor *ExecutionObject
}

type HandlersArgs struct {
	ctx          context.Context
	executionObj *ExecutionObject
	session      *wstorify.MapListStorage
	ActionsMap   *map[string]*types.ActionCardCommonType
	ImplantsMap  *map[string]*types.ActionCardCommonType
}

func (h *Handler) HandleGameplay() error {

	return nil
}

func NewHandler(executor ExecutionObject) *Handler {
	return &Handler{
		Executor: &executor,
	}
}
