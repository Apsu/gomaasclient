package api

import "github.com/apsu/gomaasclient/entity"

// NodeResults is an interface for node script results
type NodeResults interface {
	Get(params *entity.NodeResultParams) ([]entity.NodeResult, error)
}
