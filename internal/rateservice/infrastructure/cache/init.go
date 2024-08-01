package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"go_service/internal/rateservice/infrastructure"
)

func New(settings infrastructure.CacheSettings) *memcache.Client {
	return memcache.New(settings.URL)
}
