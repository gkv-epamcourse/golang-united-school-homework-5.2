package cache

import "time"

type item struct {
	key string
	val string
	deadline *time.Time
}

type Cache struct {
	items map[string]item
}

func NewCache() Cache {
	return Cache{
		items: make(map[string]item),
	}
}

func (c Cache) Get(key string) (string, bool) {
	i, ok := c.items[key]
	if !ok {
		return "", false
	}

	if i.deadline != nil {
		if i.deadline.Before(time.Now()) {
			return "", false
		}
	}

	return i.val, true
}

func (c *Cache) Put(key, value string) {
	c.items[key] = item{
		key: key,
		val: value,
	}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.items))
	for key, i := range c.items {
		if i.deadline == nil || i.deadline.After(time.Now()) {
			keys = append(keys, key)
		}		
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = item{
		key: key,
		val: value,
		deadline: &deadline,
	}
}
