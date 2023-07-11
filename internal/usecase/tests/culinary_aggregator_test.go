package usecase

import (
	"context"
	"errors"
	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"
	"github/culinary_api/internal/usecase/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := mocks.NewMockBase(ctrl)
	base.EXPECT().BeginTx(gomock.Any()).Return(context.Background(), nil).AnyTimes()
	base.EXPECT().Rollback(gomock.Any()).Return(nil).AnyTimes()
	base.EXPECT().Commit(gomock.Any()).Return(nil).AnyTimes()

	aggregatorRepo := mocks.NewMockCulinaryAggregatorRepo(ctrl)
	recipe := mocks.NewMockRecipe(ctrl)
	cookingSteps := mocks.NewMockCookingSteps(ctrl)
	ingredients := mocks.NewMockIngredients(ctrl)
	recipeIngredient := mocks.NewMockRecipeIngredient(ctrl)
	userRatings := mocks.NewMockUserRatings(ctrl)

	aggregator := usecase.NewCulinaryAggregator(base, aggregatorRepo, recipe, cookingSteps, ingredients, recipeIngredient, userRatings)

	testCases := []struct {
		userId        string
		recipeId      string
		rating        int8
		mock          func(*mocks.MockUserRatings, *mocks.MockRecipe)
		expectedError error
	}{
		{
			userId:   "d4024b9b-6a71-4b7a-86f8-f541c8a767e4",
			recipeId: "d5898e2d-1eee-4f3a-bb5b-cde925f8212d",
			rating:   5,
			mock: func(ratings *mocks.MockUserRatings, recipe *mocks.MockRecipe) {
				ratings.EXPECT().Get(gomock.Any(), map[string]string{
					"user_id":   "d4024b9b-6a71-4b7a-86f8-f541c8a767e4",
					"recipe_id": "d5898e2d-1eee-4f3a-bb5b-cde925f8212d",
				}).Return(&entity.UserRating{}, nil)
			},
			expectedError: entity.UserAlreadyVoted,
		},
		{
			userId:   "d4024b9b-6a71-4b7a-86f8-f541c8a767e4",
			recipeId: "d5898e2d-1eee-4f3a-bb5b-cde925f8212d",
			rating:   5,
			mock: func(ratings *mocks.MockUserRatings, recipe *mocks.MockRecipe) {
				userId := "d4024b9b-6a71-4b7a-86f8-f541c8a767e4"
				recipeId := "d5898e2d-1eee-4f3a-bb5b-cde925f8212d"
				rating := int8(5)

				ratings.EXPECT().Get(gomock.Any(), map[string]string{
					"user_id":   userId,
					"recipe_id": recipeId,
				}).Return(nil, entity.ErrorNotFound)

				ratings.EXPECT().Create(gomock.Any(), &entity.UserRating{
					UserID:   userId,
					RecipeID: recipeId,
					Rating:   rating,
				}).Return(nil)

				recipeEntity := &entity.Recipe{Guid: recipeId}
				recipe.EXPECT().Get(gomock.Any(), map[string]string{
					"recipe_id": recipeId,
				}).Return(recipeEntity, nil)

				recipe.EXPECT().Update(gomock.Any(), recipeEntity).Return(nil)
			},
			expectedError: nil,
		},
	}

	for _, v := range testCases {
		t.Run("Test add rating to recipe", func(t *testing.T) {
			ctx := context.Background()

			v.mock(userRatings, recipe)

			err := aggregator.AddRating(ctx, v.userId, v.recipeId, v.rating)
			if !errors.Is(err, v.expectedError) {
				t.Errorf("error user already voted FAILED")
			}
		})

	}

}
