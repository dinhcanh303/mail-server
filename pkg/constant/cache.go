package constant

import "time"

const (
	CachePrefix string = "sv:mail:"
)
const (
	CacheExpiresInOneSecond = time.Second
	CacheExpiresInOneMinute = CacheExpiresInOneSecond * 60
	CacheExpiresInOneHour   = CacheExpiresInOneMinute * 60
	CacheExpiresInOneDay    = CacheExpiresInOneHour * 24
	CacheExpiresInOneMonth  = CacheExpiresInOneDay * 30
	CacheExpiresInOneYear   = CacheExpiresInOneMonth * 12
)
