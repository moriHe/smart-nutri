package postgres

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (s *Storage) GetInvitationLink(user *types.User) (string, error) {
	var userRole string
	err := s.Db.QueryRow(
		context.Background(),
		"select user_role from users_familys where family_id = $1 and user_id = $2",
		user.ActiveFamilyId, user.Id).Scan(&userRole)

	if userRole != "OWNER" {
		return "", &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("No permition")}
	}

	token, err := generateSecureToken()
	if err != nil {
		return "", &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Could not generate token")}
	}

	currentTime := time.Now().Format(time.RFC3339)

	var query = "insert into invitations " +
		"(created_at, token, family_id) values ($1, $2, $3) returning token"

	var dbToken string
	err = s.Db.QueryRow(context.Background(), query, currentTime, token, user.ActiveFamilyId).Scan(&dbToken)

	if err != nil || err == pgx.ErrNoRows {
		return "", &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("No token found")}
	}
	return dbToken, nil

}

func addUserToFamily(db *pgxpool.Pool, userId int, token string) error {
	var familyId int

	tx, err := db.Begin(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to begin transaction"}
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select family_id from invitations where token = $1", token).Scan(&familyId)

	if err != nil || err == pgx.ErrNoRows {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Invitation not found")}
	}

	var userRole string
	err = tx.QueryRow(
		context.Background(),
		"select user_role from users_familys where family_id = $1 and user_id = $2",
		familyId, userId).Scan(&userRole)

	// If row exists, user is already part of the family
	if err != pgx.ErrNoRows {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Already part of this community")}
	}

	_, err = tx.Exec(
		context.Background(),
		"insert into users_familys (family_id, user_id, user_role) values ($1, $2, $3)",
		familyId, userId, "MEMBER")
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Could not add user to family")}
	}

	_, err = tx.Exec(context.Background(), "update users set active_family_id = $1 where users.id = $2", familyId, userId)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Could not update user active family")}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: "Failed to commit transaction"}
	}
	return nil
}

// TODO: proper rollback and delete invitation link no matter what
func (s *Storage) AcceptInvitation(userId int, token string) error {
	var queryErr error = nil
	err := addUserToFamily(s.Db, userId, token)
	if err != nil {
		queryErr = err
	}
	_, err = s.Db.Exec(context.Background(), "delete from invitations where token = $1", token)
	if err != nil {
		if queryErr != nil {
			return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Delete invitation failed")}
		}
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprint("Could not delete invitation")}
	}

	return queryErr
}
