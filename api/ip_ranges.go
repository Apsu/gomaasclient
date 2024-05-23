package api

import (
	"github.com/apsu/gomaasclient/entity"
)

// IPRanges is an interface for listing and creating IPRange records
type IPRanges interface {
	Get() ([]entity.IPRange, error)
	Create(params *entity.IPRangeParams) (*entity.IPRange, error)
}
