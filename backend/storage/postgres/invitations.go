package postgres

import (
	"context"
	"crypto/rand"
	"encoding/hex"
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
		return "", &types.InternalServerError
	}

	token, err := generateSecureToken()
	if err != nil {
		return "", &types.InternalServerError
	}

	currentTime := time.Now().Format(time.RFC3339)

	var query = "insert into invitations " +
		"(created_at, token, family_id) values ($1, $2, $3) returning token"

	var dbToken string
	err = s.Db.QueryRow(context.Background(), query, currentTime, token, user.ActiveFamilyId).Scan(&dbToken)

	if err != nil || err == pgx.ErrNoRows {
		return "", &types.InternalServerError
	}
	return dbToken, nil

}

func addUserToFamily(db *pgxpool.Pool, userId int, token string) error {
	var familyId int

	tx, err := db.Begin(context.Background())
	if err != nil {
		return &types.InternalServerError
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select family_id from invitations where token = $1", token).Scan(&familyId)

	if err != nil || err == pgx.ErrNoRows {
		return &types.InternalServerError
	}

	var userRole string
	err = tx.QueryRow(
		context.Background(),
		"select user_role from users_familys where family_id = $1 and user_id = $2",
		familyId, userId).Scan(&userRole)

	// If row exists, user is already part of the family
	if err != pgx.ErrNoRows {
		return &types.InternalServerError
	}

	_, err = tx.Exec(
		context.Background(),
		"insert into users_familys (family_id, user_id, user_role) values ($1, $2, $3)",
		familyId, userId, "MEMBER")
	if err != nil {
		return &types.InternalServerError
	}

	_, err = tx.Exec(context.Background(), "update users set active_family_id = $1 where users.id = $2", familyId, userId)
	if err != nil {
		return &types.InternalServerError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.InternalServerError
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
	_, errOnDelete := s.Db.Exec(context.Background(), "delete from invitations where token = $1", token)
	if queryErr != nil {
		return queryErr

	}
	if errOnDelete != nil {
		return &types.InternalServerError
	}

	return nil
}
