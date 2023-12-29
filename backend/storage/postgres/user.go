package postgres

import (
	"context"
	"fmt"
	"net/http"

	"github.com/moriHe/smart-nutri/types"
)

func (s *Storage) PostUser(payload types.PostUser) (*int, error) {
	var userId int

	err := s.Db.QueryRow(context.Background(), "insert into users (fire_uid) values ($1) returning id", payload.FireUid).Scan(&userId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to create user: %s", err)}
	}

	return &userId, nil
}
