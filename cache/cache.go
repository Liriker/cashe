package cache

type cache struct {
	data map[string]any
}

func New() *cache {
	data := make(map[string]any)
	return &cache{data: data}
}

func (c *cache) Set(key string, value any) {
	c.data[key] = value
}

func (c *cache) Get(key string) any {
	result := c.data[key]
	return result
}

func (c *cache) Delete(key string) {
	delete(c.data, key)
}
