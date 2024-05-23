package api

import "github.com/apsu/gomaasclient/entity"

// NodeDevice is an interface defining API behaviour for Node Devices
type NodeDevice interface {
	Get(systemID string, id int) (*entity.NodeDevice, error)
	Delete(systemID string, id int) error
}
