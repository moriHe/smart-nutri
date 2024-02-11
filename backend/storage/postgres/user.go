package postgres

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/moriHe/smart-nutri/types"
)

func marshalUser(user *types.User) (*types.User, error) {
	result, err := json.Marshal(user)
	if err != nil {
		return nil, &types.InternalServerError
	}

	var marshaledUser types.User
	err = json.Unmarshal(result, &marshaledUser)
	if err != nil {
		return nil, &types.InternalServerError
	}

	return &marshaledUser, nil
}

func (s *Storage) GetUser(fireUid string) (*types.User, error) {
	var user types.User
	err := s.Db.QueryRow(context.Background(), "select id, active_family_id from users where supabase_uid = $1", fireUid).Scan(&user.Id, &user.ActiveFamilyId)

	if err != nil {
		return nil, &types.BadRequestError
	}
	return marshalUser(&user)
}

func (s *Storage) PostUser(fireUid string) (*int, error) {
	var userId int

	err := s.Db.QueryRow(context.Background(), "insert into users (supabase_uid) values ($1) returning id", fireUid).Scan(&userId)

	if err != nil {
		return nil, &types.BadRequestError
	}

	return &userId, nil
}

func (s *Storage) PatchUser(userId int, newActiveFamilyId int) error {
	var familyId int
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.BadRequestError
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select id from users_familys where user_id = $1 and family_id = $2", userId, newActiveFamilyId).Scan(&familyId)
	if err != nil || err == pgx.ErrNoRows {
		return &types.BadRequestError
	}

	if familyId == newActiveFamilyId {
		return nil
	}

	_, err = tx.Exec(context.Background(), "update users set active_family_id = $1 where id = $2", newActiveFamilyId, userId)
	if err != nil {
		return &types.BadRequestError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.BadRequestError
	}

	return nil
}

func (s *Storage) GetUserFamilys(userId int) (*[]types.UserFamily, error) {
	rows, err := s.Db.Query(context.Background(), "select uf.id, uf.family_id, familys.name, uf.user_role from users_familys as uf join familys on uf.family_id = familys.id where uf.user_id = $1", userId)
	if err != nil {
		return nil, &types.BadRequestError
	}

	defer rows.Close()

	var familys []types.UserFamily
	for rows.Next() {
		var family types.UserFamily
		if err = rows.Scan(&family.Id, &family.FamilyId, &family.FamilyName, &family.Role); err != nil {
			return nil, &types.BadRequestError
		}
		familys = append(familys, family)
	}
	if err := rows.Err(); err != nil {
		return nil, &types.BadRequestError
	}
	return &familys, nil
}
