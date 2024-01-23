package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moriHe/smart-nutri/types"
)

func marshalUser(user *types.User) (*types.User, error) {
	result, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	var marshaledUser types.User
	err = json.Unmarshal(result, &marshaledUser)
	if err != nil {
		return nil, err
	}

	return &marshaledUser, nil
}

func (s *Storage) GetUser(fireUid string) (*types.User, error) {
	var user types.User
	err := s.Db.QueryRow(context.Background(), "select id, active_family_id from users where supabase_uid = $1", fireUid).Scan(&user.Id, &user.ActiveFamilyId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("No user: %s", err)}
	}
	return marshalUser(&user)
}

func (s *Storage) PostUser(fireUid string) (*int, error) {
	var userId int

	err := s.Db.QueryRow(context.Background(), "insert into users (supabase_uid) values ($1) returning id", fireUid).Scan(&userId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to create user: %s", err)}
	}

	return &userId, nil
}

func (s *Storage) PatchUser(userId int, newActiveFamilyId int) error {
	// todo check that userId is part of family via users_familys
	return nil
}

func (s *Storage) GetUserFamilys(userId int) (*[]types.UserFamily, error) {
	rows, err := s.Db.Query(context.Background(), "select uf.id, uf.family_id, familys.name, uf.user_role from users_familys as uf join familys on uf.family_id = familys.id where uf.user_id = $1", userId)
	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to query users_familys: %s", err)}
	}

	defer rows.Close()

	var familys []types.UserFamily
	for rows.Next() {
		fmt.Println("here")
		var family types.UserFamily
		if err = rows.Scan(&family.Id, &family.FamilyId, &family.FamilyName, &family.Role); err != nil {
			return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to scan users_familys: %s", err)}
		}
		familys = append(familys, family)
	}
	if err := rows.Err(); err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed querying users_familys: %s", err)}
	}
	return &familys, nil
}
