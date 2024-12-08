package repository

type InjectableStructs struct {
	Auth *AuthRepository
	User *UserRepository
	App *AppRepository
}


type UserCredentials struct {
	AccessToken string
	RefreshToken string
	ExpireIn int64
}