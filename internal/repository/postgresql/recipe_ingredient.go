package postgresql

import (
	"context"
	"fmt"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type recipeIngredient struct {
	db        *postgres.DB
	tableName string
}

func NewRecipeIngredientRepo(db *postgres.DB) *recipeIngredient {
	return &recipeIngredient{
		db:        db,
		tableName: "recipe_ingredients",
	}
}

func (r *recipeIngredient) Create(ctx context.Context, m *entity.RecipeIngredient) error {
	query := r.db.Builder.
		Insert(r.tableName).
		SetMap(r.getMap("create", m))

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *recipeIngredient) Update(ctx context.Context, m *entity.RecipeIngredient) error {
	query := r.db.Builder.
		Update(r.tableName).
		SetMap(r.getMap("update", m)).
		Where(r.db.Builder.Equal("recipe_id", m.RecipeId), r.db.Builder.Equal("ingredient_id", m.IngredientId))

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *recipeIngredient) FindOne(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error) {
	query := r.db.Builder.Select(
		"recipe_id",
		"ingredient_id",
		"count",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "recipe_id", "ingredient_id":
			query = query.Where(r.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result := &entity.RecipeIngredient{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.RecipeId,
		&result.IngredientId,
		&result.Count,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *recipeIngredient) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error) {
	query := r.db.Builder.Select(
		"recipe_id",
		"ingredient_id",
		"count",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "recipe_id", "ingredient_id":
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
		return nil, err
	}
	defer rows.Close()

	result := []*entity.RecipeIngredient{}
	for rows.Next() {
		temp := &entity.RecipeIngredient{}
		err = rows.Scan(
			&temp.RecipeId,
			&temp.IngredientId,
			&temp.Count,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *recipeIngredient) getMap(t string, m *entity.RecipeIngredient) map[string]interface{} {
	result := map[string]interface{}{
		"count": m.Count,
	}
	if t == "create" {
		result["recipe_id"] = m.RecipeId
		result["ingredient_id"] = m.IngredientId
	}
	return result
}

func (r *recipeIngredient) Delete(ctx context.Context, filter map[string]string) error {
	query := r.db.Builder.
		Delete(r.tableName)

	var filterExist bool
	for k, v := range filter {
		switch k {
		case "recipe_id", "ingredient_id":
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
		return err
	}
	return nil
}
