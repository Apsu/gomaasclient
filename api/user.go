package api

import "github.com/apsu/gomaasclient/entity"

// User is an interface for getting and deleting users
type User interface {
	Get(userName string) (*entity.User, error)
	Delete(userName string) error
}
