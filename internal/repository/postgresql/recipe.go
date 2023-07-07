package postgresql

import (
	"context"
	"fmt"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type recipe struct {
	db        *postgres.DB
	tableName string
}

func NewRecipeRepo(db *postgres.DB) *recipe {
	return &recipe{
		db:        db,
		tableName: "recipe",
	}
}

func (r *recipe) Create(ctx context.Context, m *entity.Recipe) error {
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

func (r *recipe) Update(ctx context.Context, m *entity.Recipe) error {
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
		return err
	}
	return nil
}

func (r *recipe) FindOne(ctx context.Context, m map[string]string) (*entity.Recipe, error) {
	query := r.db.Builder.Select(
		"guid",
		"title",
		"description",
		"created_at",
		"updated_at",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "recipe_id":
			query = query.Where(r.db.Builder.Equal("guid", v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result := &entity.Recipe{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.Guid,
		&result.Title,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *recipe) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error) {
	query := r.db.Builder.Select(
		"guid",
		"title",
		"description",
		"created_at",
		"updated_at",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid":
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

	result := []*entity.Recipe{}
	for rows.Next() {
		temp := &entity.Recipe{}
		err = rows.Scan(
			&temp.Guid,
			&temp.Title,
			&temp.Description,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *recipe) getMap(t string, m *entity.Recipe) map[string]interface{} {
	result := map[string]interface{}{
		"title":       m.Title,
		"description": m.Description,
		"updated_at":  m.UpdatedAt,
	}
	if t == "create" {
		result["guid"] = m.Guid
		result["created_at"] = m.CreatedAt

	}
	return result
}

func (r *recipe) Delete(ctx context.Context, filter map[string]string) error {
	query := r.db.Builder.Delete(r.tableName)

	var filterExist bool
	for k, v := range filter {
		switch k {
		case "guid", "recipe_id":
			query = query.Where(r.db.Builder.Equal("guid", v))
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
