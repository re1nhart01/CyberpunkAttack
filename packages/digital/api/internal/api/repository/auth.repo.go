package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/constants"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/crypto"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/postgres"
	"github.com/cyberpunkattack/database/redis"
	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/internal/external/sendpulse"
	"github.com/cyberpunkattack/jwt"
	"strings"
	"time"
)

type AuthRepository struct {
	*base.Repository
	injections struct {
		User *UserRepository
	}
}

type InitialUser struct {
	Username string
	Email string
	Phone string
	FullName string
	Password string
}


func (repo *AuthRepository) CreateInitialUser(data InitialUser) error {
	serverKey := environment.GEnv().GetVariable("SERVER_KEY")
	userHash := crypto.GetSha1(serverKey, strings.Join([]string{data.Email, data.Username, time.Now().String()}, "@"))
	hashedPassword := crypto.HashPassword(data.Password)

	newUserData := &models.UserModel{
		UserHash:    userHash,
		Username:    data.Username,
		Email:       data.Email,
		Phone:       data.Phone,
		Role:        database.ROLES_USER,
		FullName:    data.FullName,
		Temporary:   true,
		Active:      false,
		Password:    hashedPassword,
	}

	if payload := postgres.DB().Get().Table(database.USERS_TABLE).Create(&newUserData); payload.Error != nil {
		return errors.New(fmt.Sprintf(inlineErrors.REGISTER_ERROR_1, payload.Error.Error()))
	}

	return nil
}

func (repo *AuthRepository) SendEmailToRedis(email, name string) error {
	ctx := context.Background()
	inviteHash := crypto.GetSha1(environment.GEnv().Get("SERVER_KEY"), time.Now().String())
	status := redis.Db().Get().Set(ctx, fmt.Sprintf("%s:%s", constants.REDIS_INVITES_TABLE, email), inviteHash, time.Hour * 24)

	fmt.Println("inviteHash: ", inviteHash)

	if status.Err() != nil {
		return status.Err()
	}

	sp := sendpulse.New()

	user := sp.CreateUser(email, name)

	if err := sp.CreateMessage(230, constants.REGISTER_SUBJECT, constants.FROM_EMAIL, user).Send(ctx); err != nil {
		return err
	}

	return nil
}

func (repo *AuthRepository) ValidateEmail(email, code string) (bool, error) {
	ctx := context.Background()
	rdb :=  redis.
		Db().
		Get()
	key := fmt.Sprintf("%s:%s", constants.REDIS_INVITES_TABLE, email)

	redisResponse := rdb.Get(ctx, key)

	if redisResponse.Err() != nil {
		return false, redisResponse.Err()
	}

	codeFromRedis, err := redisResponse.Result()
	if err != nil {
		return false, err
	}
	fmt.Println(code, codeFromRedis)
	if !strings.EqualFold(code, codeFromRedis) {
		return false, nil
	}

	rdb.Del(ctx, key)

	if payload := postgres.
		DB().
		Get().
		Table(database.USERS_TABLE).
		Where("email = ?", email).
		Update("temporary", false); payload.Error != nil {
		return false, payload.Error
	}

	return true, nil
}


func (repo *AuthRepository) LogInUser(email, password string) error {

	user, ok := repo.injections.User.GetUserByEmail(email)

	if !ok {
		return errors.New("user with this email is not exists")
	}

	isValidPassword := crypto.CheckPasswordHash(password, user.Password)

	if !isValidPassword {
		return errors.New("password is not valid")
	}

	return nil
}


func (repo *AuthRepository) GenerateUserTokens(email string) (*UserCredentials, error) {
	user, ok := repo.injections.User.GetUserByEmail(email)

	if !ok {
		return nil, fmt.Errorf("user with email %s does not exist", email)
	}

	expiresAt := time.Now().Add(time.Hour * constants.ACCESS_TOKEN_EXPIRATION_HOURS)
	accessToken, err := jwt.CreateToken(email, user.UserHash, user.Id, constants.USER_CREDS_TOKEN_TYPE, &expiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	refreshTokenExpiration := time.Now().Add(time.Hour * constants.REFRESH_TOKEN_EXPIRATION_HOURS)
	refreshToken, err := jwt.CreateToken(email, user.UserHash, user.Id, constants.USER_CREDS_TOKEN_TYPE, &refreshTokenExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	creds := &UserCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     expiresAt.UnixMilli(),
	}

	return creds, nil
}

func (repo *AuthRepository) ValidateToken(refreshToken, grantType string) (*jwt.UserClaim, error) {
	if grantType != constants.GRANT_TYPE_REFRESH {
		return nil, errors.New("invalid grant type")
	}
	expired := jwt.CheckIsTokenExpired(refreshToken)

	if expired {
		return nil, errors.New("refresh token is expired")
	}

	claims, err := jwt.VerifyToken(refreshToken)

	if err != nil || claims.TokenType != constants.USER_CREDS_TOKEN_TYPE {
		return nil, errors.New("refresh token has invalid signature")
	}


	return claims, nil
}



func NewAuthRepository(injectable InjectableStructs) *AuthRepository {
	return &AuthRepository{
		Repository:	&base.Repository{
			TableName: "users",
			},
			injections: struct{ User *UserRepository }{User: injectable.User },
	}
}