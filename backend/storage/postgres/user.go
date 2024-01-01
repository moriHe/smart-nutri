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
	err := s.Db.QueryRow(context.Background(), "select id, active_family_id from users where fire_uid = $1", fireUid).Scan(&user.Id, &user.ActiveFamilyId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("No user: %s", err)}
	}
	return marshalUser(&user)
}

func (s *Storage) PostUser(payload types.PostUser) (*int, error) {
	var userId int

	err := s.Db.QueryRow(context.Background(), "insert into users (fire_uid) values ($1) returning id", payload.FireUid).Scan(&userId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to create user: %s", err)}
	}

	return &userId, nil
}
