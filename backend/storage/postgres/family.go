package postgres

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/moriHe/smart-nutri/types"
)

func generateUniqueKey(length int) (string, error) {
	if length%2 != 0 {
		return "", fmt.Errorf("length must be even")
	}

	// Generate random bytes
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert bytes to hexadecimal string
	randomPart := hex.EncodeToString(bytes)

	// Add timestamp to make it more unique
	timestamp := time.Now().Unix()
	uniqueKey := fmt.Sprintf("%s-%d", randomPart, timestamp)

	return uniqueKey, nil
}

func (s *Storage) PostFamily(name string, userId int) *types.RequestError {
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to begin transaction"}
	}
	defer tx.Rollback(context.Background())

	uniqueKey, err := generateUniqueKey(8)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Something went wrong"}
	}

	var familyId int
	err = tx.QueryRow(context.Background(), "insert into familys (name, code) values ($1, $2) returning id", name, uniqueKey).Scan(&familyId)

	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to create family: %s", err)}
	}

	_, err = tx.Exec(context.Background(), "update users set active_family_id = $1 where id = $2", familyId, userId)
	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to update user: %s", err)}
	}

	_, err = tx.Exec(context.Background(), "insert into users_familys (family_id, user_id, user_role) values ($1, $2, $3)", familyId, userId, "OWNER")

	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to create users_familys: %s", err)}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to commit transaction"}
	}

	return nil
}
