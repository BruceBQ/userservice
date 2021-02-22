package cache

import "time"

type CacheOptions struct {
	Size          int
	DefaultExpiry time.Duration
	Name          string
}

type Provider interface {
	NewCache(opts *CacheOptions) (Cache, error)

	Connect() error

	Close() error
}

type cacheProvider struct {
}

func NewProvider() Provider {
	return &cacheProvider{}
}

func (c *cacheProvider) NewCache(opts *CacheOptions) (Cache, error) {
	return NewLRU(LRUOptions{
		Name:          opts.Name,
		Size:          opts.Size,
		DefaultExpiry: opts.DefaultExpiry,
	}), nil
}

func (c *cacheProvider) Connect() error {
	return nil
}

func (c *cacheProvider) Close() error {
	return nil
}
