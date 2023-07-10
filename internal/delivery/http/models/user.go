package models

type RegistrationRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	RetryPassword string `json:"retry_password"`
}

type RegistrationResponse struct {
	Guid string `json:"guid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	AccessTokenTTL  int `json:"access_token_ttl"`
	RefreshTokenTTL int `json:"refresh_token_ttl"`
}
