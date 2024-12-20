package service

import (
	"context"
	"time"

	"github.com/cyberpunkattack/api/base"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/pkg/dispatcher"
)

type GatewayService struct {
	*base.Service
}

func (gateway *GatewayService) SubscribeAllEvents() *dispatcher.Dispatcher {
	dp := *dispatcher.New()
	ctx := context.Background()
	dp.AddManyListener(dispatcher.DispatcherSubscription{
		{
			Name:  "testing",
			Uname: "app.ws.global",
			Cb: func(args map[string]any) (time.Time, error) {
				mongo.DB().Get().Collection("sessions").InsertOne(ctx, &models.SessionIM{})
				return time.Now(), nil
			},
		},
	})

	return &dp
}

func (gateway *GatewayService) UnsubscribeAllEvents() error {
	return nil
}

func NewGatewayService() *GatewayService {
	return &GatewayService{
		&base.Service{},
	}
}
