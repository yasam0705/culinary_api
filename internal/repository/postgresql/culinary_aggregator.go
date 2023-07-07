package postgresql

import (
	"context"
	"github/culinary_api/internal/entity"
	"github/culinary_api/pkg/postgres"
)

type culinaryAggregator struct {
	db                                                                                       *postgres.DB
	tableNameCookingSteps, tableNameRecipe, tableNameRecipeIngredients, tableNameingredients string
}

func NewCulinaryAggregatorRepo(db *postgres.DB) *culinaryAggregator {
	return &culinaryAggregator{
		db:                         db,
		tableNameCookingSteps:      "cooking_steps",
		tableNameRecipe:            "recipe",
		tableNameRecipeIngredients: "recipe_ingredients",
		tableNameingredients:       "ingredients",
	}
}

func (c *culinaryAggregator) Ingridients(ctx context.Context, filters map[string]string) ([]*entity.Ingredients, error) {
	query := c.db.Builder.Select(
		"guid",
		"name",
		"dimension",
		"count",
	).From(c.tableNameingredients + " as i").
		Join(c.tableNameRecipeIngredients + " as ri ON ri.ingredient_id = i.guid")

	for k, v := range filters {
		switch k {
		case "guid", "ingredient_id", "recipe_id", "name":
			query = query.Where(c.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Ingredients, 0, 2)
	for rows.Next() {
		temp := &entity.Ingredients{}
		err = rows.Scan(
			&temp.Guid,
			&temp.Name,
			&temp.Dimension,
			&temp.Count,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, temp)
	}

	return list, nil
}
