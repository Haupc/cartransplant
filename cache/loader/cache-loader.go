package loader

// CacheLoader load data if not exist
type CacheLoader interface {
	Load(key string) (string, error)
}
