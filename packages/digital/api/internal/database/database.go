package database

import (
	"context"
	"sync"

	"github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/database/postgres"
)

const (
	USERS_TABLE         = "users"
	PREVIOUS_DATA_TABLE = "previous_data"
	CLANS_TABLE         = "clans"
	CLANS_MEMBER_TABLE  = "clan_member"
	BUG_REPORT_TABLE    = "bug_report"
	FRIENDS_TABLE       = "friends"
)

const (
	ROLES_USER = "User"
)

const EmailConstaint = `CONSTRAINT email_format_check CHECK (
        email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
	)`

type IModels interface {
	NewClansTable() string
	NewBugReportTable() string
	NewUserModel() string
	NewPreviousDataModel() string
	NewFriendModel() string
	NewClansMemberTable() string
}

func CreatePostgresTables(ctx context.Context, wg *sync.WaitGroup, primaryModel IModels) {
	instance := postgres.DB()
	instance.CallManualSQL(primaryModel.NewClansTable())
	instance.CallManualSQL(primaryModel.NewBugReportTable())
	instance.CallManualSQL(primaryModel.NewUserModel())
	instance.CallManualSQL(primaryModel.NewPreviousDataModel())
	instance.CallManualSQL(primaryModel.NewFriendModel())
	instance.CallManualSQL(primaryModel.NewClansMemberTable())
	defer wg.Done()
	defer ctx.Done()
}

func CreateMongoCollections(ctx context.Context, wg *sync.WaitGroup) {
	instance := mongo.DB()
	instance.CreateCollection(ctx, "sessions", 1024*1024, false)
	instance.CreateCollection(ctx, "game_queue", 1024*1024, false)
	defer wg.Done()
	defer ctx.Done()
}

func CreatePostgresFunctions(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer ctx.Done()
}
