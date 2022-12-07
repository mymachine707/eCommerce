package storage

import (
	"mymachine707/protogen/blogpost"
)

// Interfaces ...
type Interfaces interface {
	// For User
	AddUser(id string, entity *blogpost.CreateUserRequest) error
	GetUserByID(id string) (*blogpost.User, error)
	GetUserList(offset, limit int, search string) (resp *blogpost.GetUserListResponse, err error)
	UpdateUser(User *blogpost.UpdateUserRequest) error
	DeleteUser(idStr string) error

	GetUserByUsername(username string) (*blogpost.User, error)
}
