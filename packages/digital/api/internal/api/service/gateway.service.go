package service

import (
	"fmt"
	"time"

	"github.com/cyberpunkattack/api/base"
	// models "github.com/cyberpunkattack/database/model"
	// "github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/pkg/dispatcher"
)

type GatewayService struct {
	*base.Service
}

func (gateway *GatewayService) GetAllSubscriptions() dispatcher.DispatcherSubscription {
	// ctx := context.Background()
	list := dispatcher.DispatcherSubscription{
		dispatcher.NewSubscription("app.ws.global", "testing", func(args map[string]any) (time.Time, error) {
			// mongo.DB().Get().Collection("sessions").InsertOne(ctx, &models.SessionIM{})
			fmt.Println("event", args)
			return time.Now(), nil
		}),
	}

	return list
}

func (gateway *GatewayService) UnsubscribeAllEvents() error {
	return nil
}

func NewGatewayService() *GatewayService {
	return &GatewayService{
		&base.Service{},
	}
}
