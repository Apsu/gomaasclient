package api

import (
	"github.com/apsu/gomaasclient/entity"
)

// Spaces is an interface for listing and creating Space objects
type Spaces interface {
	Get() ([]entity.Space, error)
	Create(name string) (*entity.Space, error)
}
