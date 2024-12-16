package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/cyberpunkattack/environment"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDB struct {
	instance *mongo.Client
}

var mongoInstance *MongoDB

func DB() *MongoDB {
	return mongoInstance
}
func (mdb *MongoDB) Get() *mongo.Database {
	return mdb.instance.Database(environment.GEnv().Get("MONGO_DB_DATABASE"))
}

func New() *MongoDB {
	API_URL := environment.GEnv().Get("MONGO_DB_DSN")
	instance, err := mongo.Connect(options.Client().ApplyURI(API_URL))
	if err != nil {
		fmt.Println("unable to connect to db")
		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)

	defer instance.Ping(ctx, readpref.Primary())

	mongoInstance = &MongoDB{
		instance: instance,
	}

	return mongoInstance
}
