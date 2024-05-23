package api

import "github.com/apsu/gomaasclient/entity"

type Discovery interface {
	Get(id string) (*entity.Discovery, error)
}
