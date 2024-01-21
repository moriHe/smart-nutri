package postgres

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/moriHe/smart-nutri/types"
)

func generateSecureToken() (string, error) {
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the random data as a hexadecimal string
	token := hex.EncodeToString(tokenBytes)
	return token, nil
}

func (s *Storage) GetInvitationLink(familyId *int) (string, error) {
	token, err := generateSecureToken()
	if err != nil {
		return "", &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}

	currentTime := time.Now().Format(time.RFC3339)

	var query = "insert into invitations " +
		"(created_at, token, family_id) values ($1, $2, $3) returning token"

	var dbToken string
	err = s.Db.QueryRow(context.Background(), query, currentTime, token, familyId).Scan(&dbToken)

	if err != nil || dbToken == "" {
		return "", &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}
	return dbToken, nil

}

// For future: Maybe an additional step where token includes username and it can be checked if users_familys has
// it listed as PROSPECTIVE
func (s *Storage) AcceptInvitation(userId int, token string) error {
	var familyId int
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to begin transaction"}
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select family_id from invitations where token = $1", token).Scan(&familyId)

	if err != nil || familyId == 0 {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}

	var userRole string
	err = tx.QueryRow(
		context.Background(),
		"select user_role from users_familys where family_id = $1 and user_id = $2",
		familyId, userId).Scan(&userRole)

	if err != sql.ErrNoRows || err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Already part of this community")}
	}

	_, err = tx.Exec(
		context.Background(),
		"insert into users_familys (family_id, user_id, user_role) values ($1, $2, $3)",
		familyId, userId, "MEMBER")
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}

	_, err = tx.Exec(context.Background(), "update users set active_family_id = $1 where users.id = $2", familyId, userId)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}

	_, err = tx.Exec(context.Background(), "delete from invitations where token = $1", token)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Something went wrong")}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to commit transaction"}
	}

	return nil
}
