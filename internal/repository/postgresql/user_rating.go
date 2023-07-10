package postgresql

import (
	"context"

	"github/culinary_api/internal/entity"

	"github/culinary_api/pkg/postgres"
)

type userRatings struct {
	db        *postgres.DB
	tableName string
}

func NewUserRatingsRepo(db *postgres.DB) *userRatings {
	return &userRatings{
		db:        db,
		tableName: "user_ratings",
	}
}

func (r *userRatings) Create(ctx context.Context, m *entity.UserRating) error {
	query := r.db.Builder.
		Insert(r.tableName).
		SetMap(r.getMap("create", m))

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return r.db.PgErr(err)
	}
	return nil
}

func (r *userRatings) Update(ctx context.Context, m *entity.UserRating) error {
	query := r.db.Builder.
		Update(r.tableName).
		SetMap(r.getMap("update", m)).
		Where(r.db.Builder.Equal("guid", m.Guid))

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return r.db.PgErr(err)
	}
	return nil
}

func (r *userRatings) FindOne(ctx context.Context, m map[string]string) (*entity.UserRating, error) {
	query := r.db.Builder.Select(
		"guid",
		"user_id",
		"recipe_id",
		"rating",
		"created_at",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid", "user_id", "recipe_id":
			query = query.Where(r.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result := &entity.UserRating{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.Guid,
		&result.UserID,
		&result.RecipeID,
		&result.Rating,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, r.db.PgErr(err)
	}

	return result, nil
}

func (r *userRatings) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.UserRating, error) {
	query := r.db.Builder.Select(
		"guid",
		"user_id",
		"recipe_id",
		"rating",
		"created_at",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid", "user_id":
			query = query.Where(r.db.Builder.Equal(k, v))
		}
	}

	if limit != 0 {
		query = query.Limit(limit).Offset(offset)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, r.db.PgErr(err)
	}
	defer rows.Close()

	result := []*entity.UserRating{}
	for rows.Next() {
		temp := &entity.UserRating{}
		err = rows.Scan(
			&temp.Guid,
			&temp.UserID,
			&temp.RecipeID,
			&temp.Rating,
			&temp.CreatedAt,
		)
		if err != nil {
			return nil, r.db.PgErr(err)
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *userRatings) getMap(t string, m *entity.UserRating) map[string]interface{} {
	result := map[string]interface{}{
		"rating": m.Rating,
	}
	if t == "create" {
		result["guid"] = m.Guid
		result["user_id"] = m.UserID
		result["recipe_id"] = m.RecipeID
		result["created_at"] = m.CreatedAt
	}
	return result
}
