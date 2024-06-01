package storagecache

import (
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

var StorageCache = map[schema.GroupResource]rest.Storage{}
var lock sync.RWMutex

func Add(gr schema.GroupResource, s rest.Storage) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := StorageCache[gr]; ok {
		return
	}
	StorageCache[gr] = s
}

func Get(gr schema.GroupResource) rest.Storage {
	lock.RLock()
	defer lock.RUnlock()
	if res, ok := StorageCache[gr]; ok {
		return res
	}
	return nil
}
