package handlers

import (
	"fmt"

	"github.com/cyberpunkattack/api/base"
)

const GATEWAY_ROUTE = "gateway"

type IGatewayRepo interface {
}

type GatewayHandler struct {
	*base.Handler
	IGatewayRepo
}

func (gateway *GatewayHandler) GetName() string {
	return gateway.Name
}

func (gateway *GatewayHandler) GetPath() string {
	return gateway.Path
}

func NewGatewayHandler(basePath string, repo IUserRepo) *GatewayHandler {
	return &GatewayHandler{
		&base.Handler{
			Name: GATEWAY_ROUTE,
			Path: fmt.Sprintf("/%s/%s", basePath, GATEWAY_ROUTE),
		},
		repo,
	}
}
