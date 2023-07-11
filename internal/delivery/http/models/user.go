package models

type RegistrationRequest struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	RetryPassword string `json:"retry_password" validate:"required"`
}

type RegistrationResponse struct {
	Guid string `json:"guid"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	AccessTokenTTL  int `json:"access_token_ttl"`
	RefreshTokenTTL int `json:"refresh_token_ttl"`
}
