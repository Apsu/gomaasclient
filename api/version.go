package api

import (
	"github.com/apsu/gomaasclient/entity"
)

// Version is an interface for listing MAAS version details
type Version interface {
	Get() (*entity.Version, error)
}
