package usecase

import (
	"context"
	"time"

	"github/culinary_api/internal/entity"

	"github.com/google/uuid"
)

type UserRatingsRepo interface {
	Create(ctx context.Context, userRating *entity.UserRating) error
	Update(ctx context.Context, userRating *entity.UserRating) error
	FindOne(ctx context.Context, filter map[string]string) (*entity.UserRating, error)
	FindAll(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.UserRating, error)
}

//go:generate mockgen -destination=tests/mocks/user_rating.go -package=mocks -source=user_rating.go
type UserRatings interface {
	Create(ctx context.Context, userRating *entity.UserRating) error
	Update(ctx context.Context, userRating *entity.UserRating) error
	Get(ctx context.Context, filter map[string]string) (*entity.UserRating, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.UserRating, error)
}

type userRatings struct {
	repo UserRatingsRepo
}

func NewUserRatings(repo UserRatingsRepo) *userRatings {
	return &userRatings{
		repo: repo,
	}
}

func (uc *userRatings) beforeCreate(userRating *entity.UserRating) {
	userRating.Guid = uuid.New().String()
	userRating.CreatedAt = time.Now()
}

func (uc *userRatings) Create(ctx context.Context, userRating *entity.UserRating) error {
	uc.beforeCreate(userRating)
	return uc.repo.Create(ctx, userRating)
}

func (uc *userRatings) Update(ctx context.Context, userRating *entity.UserRating) error {
	return uc.repo.Update(ctx, userRating)
}

func (uc *userRatings) Get(ctx context.Context, filter map[string]string) (*entity.UserRating, error) {
	return uc.repo.FindOne(ctx, filter)
}

func (uc *userRatings) List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.UserRating, error) {
	return uc.repo.FindAll(ctx, limit, offset, filter)
}
