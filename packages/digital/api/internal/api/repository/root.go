package repository

type InjectableStructs struct {
	Auth *AuthRepository
	User *UserRepository
	App *AppRepository
}


type UserCredentials struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireIn int64 `json:"expireIn"`
}