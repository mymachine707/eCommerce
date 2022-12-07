package postgres

import (
	"errors"
	"fmt"
	"time"

	// "mymachine707/blogpost"
	"mymachine707/protogen/blogpost"
)

// AddUser ...
func (stg Postgres) AddUser(id string, entity *blogpost.CreateUserRequest) error {
	_, err := stg.db.Exec(`INSERT INTO "user" (
		id,
		username,
		password,
		user_type
		) VALUES (
		$1,
		$2,
		$3,
		$4
)`,
		id,
		entity.Username,
		entity.Password,
		entity.UserType,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID ...  //  ????
func (stg Postgres) GetUserByID(id string) (*blogpost.User, error) {
	// var res blogpost.GetUserByIDResponse
	res := &blogpost.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.db.QueryRow(`SELECT 
    id,
	username,
	password,
	user_type,
	created_at,
	updated_at,
	deleted_at
    FROM "user" WHERE id = $1`, id).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&res.UserType,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if deletedAt != nil {
		return res, errors.New("user not found")
	}

	return res, err
}

// GetUserList ...
func (stg Postgres) GetUserList(offset, limit int, search string) (*blogpost.GetUserListResponse, error) {

	resp := &blogpost.GetUserListResponse{
		Users: make([]*blogpost.User, 0),
	}

	rows, err := stg.db.Queryx(`
	SELECT 
	id,
	username,
	password,
	user_type,
	created_at,
	updated_at
	FROM "user" WHERE deleted_at is null AND
	(username ILIKE '%' || $1 || '%') 
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &blogpost.User{}

		var updatedAt *string
		err := rows.Scan(
			&a.Id,
			&a.Username,
			&a.Password,
			&a.UserType,
			&a.CreatedAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Users = append(resp.Users, a)
	}

	return resp, err
}

// UpdateUser ...
func (stg Postgres) UpdateUser(entity *blogpost.UpdateUserRequest) error {

	res, err := stg.db.NamedExec(`UPDATE "user" SET password=:p, updated_at=now() WHERE deleted_at is null AND id=:id`, map[string]interface{}{
		"id": entity.Id,
		"p":  entity.Password,
	})

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("user not found")
}

// DeleteUser ...
func (stg Postgres) DeleteUser(idStr string) error {

	res, err := stg.db.Exec(`UPDATE "user" Set deleted_at=now() WHERE id=$1 AND deleted_at is null`, idStr)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("user not found")
}

// GetUserByUsername ...  //  ????
func (stg Postgres) GetUserByUsername(username string) (*blogpost.User, error) {
	// var res blogpost.GetUserByIDResponse
	res := &blogpost.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.db.QueryRow(`SELECT 
    	id,
		username,
		password,
		user_type,
		created_at,
		updated_at,
		deleted_at
    FROM "user" WHERE username = $1`, username).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&res.UserType,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if deletedAt != nil {
		fmt.Printf("deleted at: %s", deletedAt)
		return res, errors.New("user not found")
	}

	return res, err
}
