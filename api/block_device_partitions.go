package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type BlockDevicePartitions interface {
	Get(systemID string, blockDeviceID int) ([]entity.BlockDevicePartition, error)
	Create(systemID string, blockDeviceID int, params *entity.BlockDevicePartitionParams) (*entity.BlockDevicePartition, error)
}
