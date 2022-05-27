package cache

import "time"

type Cache struct {
	CacheMap map[string]ValueWithDeadline
}

type ValueWithDeadline struct {
	Value string
	Deadline *time.Time
}
func NewCache() Cache {
	cacheMap := make(map[string]ValueWithDeadline)
	return Cache{CacheMap: cacheMap}
}

func (c Cache) Get(key string) (string, bool) {
	result, ok := c.CacheMap[key]
	if !ok {
		return "", false
	}
	if result.Deadline != nil &&  result.Deadline.Before(time.Now()) {
		delete(c.CacheMap, key)
		return "", false
	}
	return result.Value, true
}

func (c Cache) Put(key, value string) {
	c.CacheMap[key] = ValueWithDeadline{
		Value:    value,
	}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.CacheMap))
	for k, v := range c.CacheMap {
		if v.Deadline != nil && v.Deadline.Before(time.Now()) {
			continue
		}
			keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.CacheMap[key] = ValueWithDeadline{
		Value:    value,
		Deadline: &deadline,
	}
}
