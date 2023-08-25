package postgres

import (
	"context"
	"fmt"
	"net/http"

	"github.com/moriHe/smart-nutri/types"
)

func (s *Storage) GetMealPlan(familyId string, date string) (*[]types.ShallowMealPlanItem, error) {
	query := "select mealplans.id, recipes.name, cast(date as text), meals.meal from mealplans " +
		"join recipes on mealplans.recipe_id = recipes.id join meals on mealplans.meal = meals.id " +
		"where mealplans.family_id = $1 and mealplans.date = $2"
	rows, _ := s.Db.Query(context.Background(), query, familyId, date)
	defer rows.Close()

	mealPlan := []types.ShallowMealPlanItem{}

	for rows.Next() {
		var mealPlanItem types.ShallowMealPlanItem
		err := rows.Scan(&mealPlanItem.Id, &mealPlanItem.RecipeName, &mealPlanItem.Date, &mealPlanItem.Meal)

		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan Get mealPlanItems failed: %s", err)}
		}
		mealPlan = append(mealPlan, mealPlanItem)
	}

	return &mealPlan, nil
}

func (s *Storage) GetMealPlanItem(id string) (*types.FullMealPlanItem, error) {
	return nil, nil
}

func (s *Storage) PostMealPlanItem(familyId string, payload types.PostMealPlanItem) error {
	return nil
}

func (s *Storage) PatchMealPlanItem(id string, payload types.PatchMealPlanItem) error {
	return nil
}

func (s *Storage) DeleteMealPlanItem(id string) error {
	return nil
}
