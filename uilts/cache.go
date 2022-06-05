package uilt

// 实现一个cache缓存，使用Window-Tiny-LRU算法作为缓存策略

type Cache struct {
	// 缓存的最大容量
	MaxSize int
	// 缓存的最大层数
	MaxLevel int
	// 缓存的最小层数
	MinLevel int
}


// 创建一个新的cache
func NewCache(maxSize int, maxLevel int, minLevel int) *Cache {
	return &Cache{
		MaxSize: maxSize,
		MaxLevel: maxLevel,
		MinLevel: minLevel,
	}
}

