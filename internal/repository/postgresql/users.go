package postgresql

import (
	"context"
	"fmt"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type users struct {
	db        *postgres.DB
	tableName string
}

func NewUsersRepo(db *postgres.DB) *users {
	return &users{
		db:        db,
		tableName: "users",
	}
}

func (r *users) Create(ctx context.Context, m *entity.User) error {
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

func (r *users) Update(ctx context.Context, m *entity.User) error {
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

func (r *users) FindOne(ctx context.Context, m map[string]string) (*entity.User, error) {
	query := r.db.Builder.Select(
		"guid",
		"username",
		"password",
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

	result := &entity.User{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&result.Guid,
		&result.Username,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *users) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.User, error) {
	query := r.db.Builder.Select(
		"guid",
		"username",
		"password",
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

	result := []*entity.User{}
	for rows.Next() {
		temp := &entity.User{}
		err = rows.Scan(
			&temp.Guid,
			&temp.Username,
			&temp.Password,
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

func (r *users) getMap(t string, m *entity.User) map[string]interface{} {
	result := map[string]interface{}{
		"password":   m.Password,
		"updated_at": m.UpdatedAt,
	}
	if t == "create" {
		result["guid"] = m.Guid
		result["username"] = m.Username
		result["created_at"] = m.CreatedAt

	}
	return result
}

func (r *users) Delete(ctx context.Context, filter map[string]string) error {
	query := r.db.Builder.Delete(r.tableName)

	var filterExist bool
	for k, v := range filter {
		switch k {
		case "guid":
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
