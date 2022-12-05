package cash

type cache struct {
	data map[string]any
}

func (c *cache) New() *cache {
	var data map[string]any
	return &cache{data: data}
}

func (c *cache) Set(key string) {

}

func (c *cache) Get(key string) int {
	var result int
	return result
}

func (c *cache) Delete(key string) {

}
