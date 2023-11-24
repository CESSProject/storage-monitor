package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CacheItem struct {
	FlashTime time.Time
	Data      []MinerInfoDisplay
}

type DataCache map[string]CacheItem

type CacheHandler interface {
	GetCacheData() []MinerInfoDisplay
	UpdateCache(key string, value []MinerInfoDisplay)
}

const MAX_DATA_SIZE = 100 * 100

var cache DataCache

func InitCache() {
	if cache == nil {
		cache = make(DataCache)
	}
}

func GetCacher() CacheHandler {
	return cache
}

func (c DataCache) GetCacheData() []MinerInfoDisplay {
	var data []MinerInfoDisplay

	for k, v := range c {
		if time.Since(v.FlashTime) > time.Minute*2 {
			continue
		}
		for _, item := range v.Data {
			item.ContainerInfo.Name += " " + k
			data = append(data, item)
		}
	}
	return data
}

func (c DataCache) UpdateCache(key string, value []MinerInfoDisplay) {
	if len(c) >= MAX_DATA_SIZE {
		return
	}
	c[key] = CacheItem{
		FlashTime: time.Now(),
		Data:      value,
	}
}

func GetCacheData(c *gin.Context) {
	data := cache.GetCacheData()
	c.JSON(200, data)
}

func ReceivedCacheData(c *gin.Context) {
	var data []MinerInfoDisplay
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	}
	if len(data) > 0 {
		cache.UpdateCache(c.RemoteIP(), data)
	}
	c.JSON(200, gin.H{"status": "ok"})
}
