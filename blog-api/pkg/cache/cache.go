package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCache(rdb *redis.Client, ctx context.Context) *Cache {
	return &Cache{rdb: rdb, ctx: ctx}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration, tags []string) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// store main key
	if err := c.rdb.Set(c.ctx, key, data, ttl).Err(); err != nil {
		return err
	}

	// store tag references
	for _, tag := range tags {
		tagKey := "tag:" + tag
		c.rdb.SAdd(c.ctx, tagKey, key)
	}

	return nil
}

func (c *Cache) Get(key string, dest interface{}) (bool, error) {
	val, err := c.rdb.Get(c.ctx, key).Result()
	if err != nil {
		return false, nil
	}

	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Cache) InvalidateTag(tag string) error {
	tagKey := "tag:" + tag

	keys, err := c.rdb.SMembers(c.ctx, tagKey).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		c.rdb.Del(c.ctx, keys...)
	}

	// delete tag set itself
	c.rdb.Del(c.ctx, tagKey)

	return nil
}
