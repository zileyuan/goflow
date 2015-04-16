package entity

import (
	"goflow/model"
)

type Cache map[string]model.ProcessModel

type CacheManager struct {
	Caches map[string]Cache
}
