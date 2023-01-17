package db

import (
	"errors"
	"strings"
)

var (
	// ErrNotFindTargetURL ...
	ErrNotFindTargetURL = errors.New("ErrNotFindTargetURL")
	// ErrNotFindServerAddr ...
	ErrNotFindServerAddr = errors.New("ErrNotFindServerAddr")
)

// InWhitelist 查询url是否在白名单中
func InWhitelist(url string) bool {
	synclock.RLock()
	defer synclock.RUnlock()
	_, ok := cacheURLWhiteList[url]
	return ok
}

// ParseToken 解析token
func ParseToken(token string) (username string, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	return
}

// HavePermission 验证username是否有权限
func HavePermission(username, url string) bool {
	synclock.RLock()
	defer synclock.RUnlock()
	groups, ok := cacheUserGroup[username]
	if !ok {
		return false
	}
	for _, group := range groups {
		if permission, ok1 := cacheGroupURL[group]; ok1 {
			if _, ok2 := permission[url]; ok2 {
				return true
			}
		}
	}
	return false
}

// GetEndPoint ...
func GetEndPoint(originURL string) (serverAddr []string, targetURL string, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	serverName := ""
	s, ok := cacheURL2EndPoint[originURL]
	if !ok {
		err = ErrNotFindTargetURL
		return
	}
	idx := strings.IndexByte(s, '/')
	if idx == -1 {
		serverName = s
		targetURL = ""
	} else {
		serverName = s[0:idx]
		targetURL = s[idx:]
	}
	serverAddr, ok = cacheSvcName2SvcAddr[serverName]
	if !ok {
		err = ErrNotFindServerAddr
		return
	}
	return
}
