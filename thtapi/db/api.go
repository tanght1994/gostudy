package db

import (
	"errors"
	"math/rand"
)

var (
	// ErrNotFindTargetURL ...
	ErrNotFindTargetURL = errors.New("ErrNotFindTargetURL")
	// ErrNotFindServerAddr ...
	ErrNotFindServerAddr = errors.New("ErrNotFindServerAddr")
	// ErrNoSlashInEndpoint ...
	ErrNoSlashInEndpoint = errors.New("ErrNoSlashInEndpoint")
	// ErrNotFindEndpoint ...
	ErrNotFindEndpoint = errors.New("ErrNotFindEndpoint")
)

// ParseToken 解析token
func ParseToken(token string) (username string, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	return
}

// GetURLProxy ...
func GetURLProxy(originURL string) (endpoint string, accessGroup []uint32, noPermissionCheck bool, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	oneURLProxy, ok := cacheProxyConfig[originURL]
	if !ok {
		err = ErrNotFindEndpoint
		return
	}
	if len(oneURLProxy.Endpoint) == 0 {
		err = ErrNotFindEndpoint
	}
	endpoint = oneURLProxy.Endpoint[rand.Intn(len(oneURLProxy.Endpoint))]
	accessGroup = oneURLProxy.AccessGroup
	noPermissionCheck = oneURLProxy.NoPermissionCheck
	return
}

// GetUserGroups 获取username的Group
func GetUserGroups(username string) []uint32 {
	return cacheUserGroup[username]
}
