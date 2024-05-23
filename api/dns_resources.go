package api

import (
	"github.com/apsu/gomaasclient/entity"
)

// DNSResources is an interface for listing and creating
// DNSResource records
type DNSResources interface {
	Get() ([]entity.DNSResource, error)
	Create(params *entity.DNSResourceParams) (*entity.DNSResource, error)
}
