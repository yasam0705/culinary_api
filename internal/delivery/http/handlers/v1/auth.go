package handlers

import (
	"fmt"
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/helper"
	"github/culinary_api/internal/delivery/http/models"
	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	refreshTokenType = "refresh_token"
	accessTokenType  = "access_token"
)

type authHandlers struct {
	secret                          string
	authUseCase                     usecase.Auth
	accessTokenTTL, refreshTokenTTL time.Duration
}

func NewAuthHandlers(e *gin.RouterGroup, secret string, accessTokenTTL, refreshTokenTTL time.Duration, authUseCase usecase.Auth, middleware ...gin.HandlerFunc) {
	h := authHandlers{

		secret:          secret,
		authUseCase:     authUseCase,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}

	auth := e.Group("/auth", middleware...)
	{
		auth.POST("/registration", h.Registration)
		auth.POST("/login", h.Login)
	}
}

// @Security ApiKeyAuth
// @Router /v1/auth/registration [POST]
// @Summary Registration user
// @Description Registration user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.RegistrationRequest true "data"
// @Success 200 {object} models.RegistrationResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *authHandlers) Registration(c *gin.Context) {
	ctx := c.Request.Context()

	reqBody := &models.RegistrationRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	if reqBody.Password != reqBody.RetryPassword {
		errors_pkg.Error(c, fmt.Errorf("password mismatch"))
		return
	}

	i := &entity.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	if err := r.authUseCase.Registration(ctx, i); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.CreateStepResponse{
		Guid: i.Guid,
	})
}

// @Security ApiKeyAuth
// @Router /v1/auth/login [POST]
// @Summary Registration user
// @Description Registration user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "data"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *authHandlers) Login(c *gin.Context) {
	ctx := c.Request.Context()

	reqBody := &models.LoginRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	i := &entity.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
	}
	user, err := r.authUseCase.Login(ctx, i)
	if err != nil {
		errors_pkg.Error(c, err)
		return
	}

	// jwt
	accessToken, err := helper.CreateToken(user.Guid, r.secret, r.accessTokenTTL, map[string]string{
		"typ": accessTokenType,
	})
	if err != nil {
		errors_pkg.Error(c, err)
		return
	}

	refreshToken, err := helper.CreateToken(user.Guid, r.secret, r.refreshTokenTTL, map[string]string{
		"typ": accessTokenType,
	})
	if err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.LoginResponse{
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		AccessTokenTTL:  int(r.accessTokenTTL.Seconds()),
		RefreshTokenTTL: int(r.refreshTokenTTL.Seconds()),
	})
}
