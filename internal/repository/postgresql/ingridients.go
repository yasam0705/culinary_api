package postgresql

import (
	"context"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type ingredients struct {
	db        *postgres.DB
	tableName string
}

func NewIngredientsRepo(db *postgres.DB) *ingredients {
	return &ingredients{
		db:        db,
		tableName: "ingredients",
	}
}

func (r *ingredients) Create(ctx context.Context, m *entity.Ingredients) error {
	query := r.db.Builder.
		Insert(r.tableName).
		SetMap(r.getMap("create", m)).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING guid")

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	var guid string
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&guid,
	)
	if err != nil {
		return r.db.PgErr(err)
	}
	m.Guid = guid
	return nil
}

func (r *ingredients) Update(ctx context.Context, m *entity.Ingredients) error {
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

func (r *ingredients) FindOne(ctx context.Context, m map[string]string) (*entity.Ingredients, error) {
	query := r.db.Builder.Select(
		"guid",
		"name",
		"dimension",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid", "name":
			query = query.Where(r.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result := &entity.Ingredients{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.Guid,
		&result.Name,
		&result.Dimension,
	)
	if err != nil {
		return nil, r.db.PgErr(err)
	}

	return result, nil
}

func (r *ingredients) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Ingredients, error) {
	query := r.db.Builder.Select(
		"guid",
		"name",
		"dimension",
	).From(r.tableName)

	for k, v := range m {
		switch k {
		case "guid", "name":
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

	result := []*entity.Ingredients{}
	for rows.Next() {
		temp := &entity.Ingredients{}
		err = rows.Scan(
			&temp.Guid,
			&temp.Name,
			&temp.Dimension,
		)
		if err != nil {
			return nil, r.db.PgErr(err)
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *ingredients) getMap(t string, m *entity.Ingredients) map[string]interface{} {
	result := map[string]interface{}{
		"name":      m.Name,
		"dimension": m.Dimension,
	}
	if t == "create" {
		result["guid"] = m.Guid

	}
	return result
}
