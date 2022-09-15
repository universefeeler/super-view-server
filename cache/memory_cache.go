package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	MemoryCache *cache.Cache
)

func init() {
	MemoryCache = cache.New(6*time.Hour, 12*time.Hour)
}
