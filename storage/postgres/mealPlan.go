package postgres

import "github.com/moriHe/smart-nutri/types"

func (s *Storage) GetMealPlan(familyId string, date string) (*[]types.ShallowMealPlanItem, error) {
	return nil, nil
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
