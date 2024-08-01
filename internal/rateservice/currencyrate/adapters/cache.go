package adapters

import (
	"encoding/binary"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"math"
	"time"
)

type CacheRateAdapter struct {
	client *memcache.Client
}

func NewCacheRateAdapter(client *memcache.Client) CacheRateAdapter {
	return CacheRateAdapter{client: client}
}

func (reader CacheRateAdapter) GetCurrencyRate(from string, to string) (float32, error) {
	item, err := reader.client.Get(from + ":" + to)
	if err != nil {
		return 0, err
	}
	return bytesToFloat32(item.Value)
}

func (reader *CacheRateAdapter) SetCurrencyRate(from string, to string, rate float32) error {
	return reader.client.Set(&memcache.Item{
		Key:        from + ":" + to,
		Value:      float32ToBytes(rate),
		Expiration: secondsUntilEndOfDay(),
	})
}

func bytesToFloat32(b []byte) (float32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("invalid byte length for float32")
	}
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits), nil
}

func float32ToBytes(f float32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, math.Float32bits(f))
	return b
}

func secondsUntilEndOfDay() int32 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
	return int32(endOfDay.Sub(now).Seconds())
}
