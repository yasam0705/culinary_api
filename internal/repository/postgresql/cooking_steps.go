package postgresql

import (
	"context"
	"fmt"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type cookingSteps struct {
	db        *postgres.DB
	tableName string
}

func NewCookingStepsRepo(db *postgres.DB) *cookingSteps {
	return &cookingSteps{
		db:        db,
		tableName: "cooking_steps",
	}
}

func (r *cookingSteps) Create(ctx context.Context, m *entity.CookingSteps) error {
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

func (r *cookingSteps) Update(ctx context.Context, m *entity.CookingSteps) error {
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

func (r *cookingSteps) FindOne(ctx context.Context, m map[string]string) (*entity.CookingSteps, error) {
	query := r.db.Builder.Select(
		"guid",
		"recipe_id",
		"order_number",
		"description",
		"cooking_time",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid":
			query = query.Where(r.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result := &entity.CookingSteps{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.Guid,
		&result.RecipeId,
		&result.OrderNumber,
		&result.Description,
		&result.CookingTime,
	)
	if err != nil {
		return nil, r.db.PgErr(err)
	}

	return result, nil
}

func (r *cookingSteps) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.CookingSteps, error) {
	query := r.db.Builder.Select(
		"guid",
		"recipe_id",
		"order_number",
		"description",
		"cooking_time",
	).From(r.tableName).OrderBy("order_number ASC")

	for k, v := range m {
		switch k {
		case "guid", "recipe_id":
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

	result := []*entity.CookingSteps{}
	for rows.Next() {
		temp := &entity.CookingSteps{}
		err = rows.Scan(
			&temp.Guid,
			&temp.RecipeId,
			&temp.OrderNumber,
			&temp.Description,
			&temp.CookingTime,
		)
		if err != nil {
			return nil, r.db.PgErr(err)
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *cookingSteps) getMap(t string, m *entity.CookingSteps) map[string]interface{} {
	result := map[string]interface{}{
		"order_number": m.OrderNumber,
		"description":  m.Description,
		"cooking_time": m.CookingTime,
	}
	if t == "create" {
		result["guid"] = m.Guid
		result["recipe_id"] = m.RecipeId
	}
	return result
}

func (r *cookingSteps) Delete(ctx context.Context, filter map[string]string) error {
	query := r.db.Builder.
		Delete(r.tableName)

	var filterExist bool
	for k, v := range filter {
		switch k {
		case "guid", "recipe_id":
			query = query.Where(r.db.Builder.Equal(k, v))
			filterExist = true
		}
	}
	if !filterExist {
		return fmt.Errorf("filter not exist")
	}

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
