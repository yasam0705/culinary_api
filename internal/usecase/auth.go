package usecase

import (
	"context"
	"errors"
	"fmt"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/hasher"
)

//go:generate mockgen -destination=tests/mocks/auth.go -package=mocks -source=auth.go
type Auth interface {
	Registration(ctx context.Context, m *entity.User) error
	Login(ctx context.Context, m *entity.User) (*entity.User, error)
}

type auth struct {
	userUseCase User
}

func NewAuth(userUseCase User) *auth {
	return &auth{
		userUseCase: userUseCase,
	}
}

func (a *auth) beforeRegistration(m *entity.User) (err error) {
	m.Password, err = hasher.Generate(m.Password)
	return err
}

func (a *auth) Registration(ctx context.Context, m *entity.User) error {
	existUser, err := a.userUseCase.Get(ctx, map[string]string{
		"username": m.Username,
	})
	if err != nil && !errors.Is(err, entity.ErrorNotFound) {
		return err
	}
	if existUser != nil {
		return fmt.Errorf("username already register")
	}

	err = a.beforeRegistration(m)
	if err != nil {
		return err
	}
	return a.userUseCase.Create(ctx, m)
}

func (a *auth) Login(ctx context.Context, m *entity.User) (*entity.User, error) {
	existUser, err := a.userUseCase.Get(ctx, map[string]string{
		"username": m.Username,
	})
	if err != nil {
		return nil, err
	}

	if err = hasher.Compare(existUser.Password, m.Password); err != nil {
		return nil, err
	}
	return existUser, nil
}
