package postgres

import (
	"context"
	"encoding/json"
	"fmt"

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

	err := s.Db.QueryRow(context.Background(), "insert into users (supabase_uid) values ($1) on conflict do nothing returning id", fireUid).Scan(&userId)

	if err != nil {
		return nil, &types.BadRequestError
	}

	return &userId, nil
}

type AffiliatedFamily struct {
	Id       int
	Role     string
	FamilyId int
}

func (s *Storage) DeleteUser(userId int) error {
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}
	defer tx.Rollback(context.Background())

	// set active family id to null to avoid foreign key conflicts
	_, err = tx.Exec(context.Background(), "update users set active_family_id = NULL where id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}

	// select all familys the user is affiliated with
	rows, err := tx.Query(context.Background(), "select id, family_id, user_role from users_familys where user_id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}

	affiliatedFamilys := make([]AffiliatedFamily, 0)

	for rows.Next() {
		var aFamily AffiliatedFamily
		err := rows.Scan(&aFamily.Id, &aFamily.FamilyId, &aFamily.Role)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}
		// If user is not the owner, they can be deleted from users_familys
		if aFamily.Role != "OWNER" {
			err := tx.QueryRow(context.Background(), "delete from users_familys where id = $1", aFamily.Id)
			if err != nil {
				fmt.Println(err)
				return &types.InternalServerError
			}
		} else {
			// If user is the owner, further actions need to be taken
			affiliatedFamilys = append(affiliatedFamilys, aFamily)
		}
	}

	wholeDeletion := make([]int, 0)
	for _, aFamily := range affiliatedFamilys {
		var id int
		// Check if there is at least one other member in the user's (OWNER) family
		err := tx.QueryRow(context.Background(), "select id from users_familys where family_id = $1 and user_id != $2", aFamily.FamilyId, userId).Scan(&id)
		if err != nil {
			if err == pgx.ErrNoRows {
				// If no other members, further action to delete the user is processed in the next loop
				wholeDeletion = append(wholeDeletion, aFamily.FamilyId)
			} else {
				fmt.Println(err)
				return &types.InternalServerError
			}
		} else {
			// If there is another member, the other member gets assigned the user_role OWNER
			_, err = tx.Exec(context.Background(), "update users_familys set user_role = $1 where id = $2", "OWNER", id)
			if err != nil {
				fmt.Println(err)
				return &types.InternalServerError
			}
			// Any pending invitation links generated from the old OWNER can be deleted
			_, err = tx.Exec(context.Background(), "delete from invitations where family_id = $1", aFamily.Id)
			if err != nil {
				fmt.Println(err)
				return &types.InternalServerError
			}
		}
	}

	// This runs the delete logic for each instance where the user was the sole person in the family
	for _, familyId := range wholeDeletion {
		_, err = tx.Exec(context.Background(), "DELETE FROM invitations WHERE family_id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}

		_, err = tx.Exec(context.Background(), "DELETE FROM shopping_list WHERE family_id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}

		_, err = tx.Exec(context.Background(), "DELETE FROM mealplans WHERE family_id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}
		_, err = tx.Exec(context.Background(), "DELETE FROM recipes WHERE family_id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}
		_, err = tx.Exec(context.Background(), "DELETE FROM users_familys WHERE family_id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}

		_, err = tx.Exec(context.Background(), "DELETE FROM familys WHERE id = $1", familyId)
		if err != nil {
			fmt.Println(err)
			return &types.InternalServerError
		}
	}

	// Lastly, delete the user from the final tables users_familys and users.
	// This is generally a further action to "If there is another member, the other member gets assigned the user_role OWNER"
	_, err = tx.Exec(context.Background(), "delete from users_familys where user_id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}

	_, err = tx.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Println(err)
		return &types.InternalServerError
	}

	return nil

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
