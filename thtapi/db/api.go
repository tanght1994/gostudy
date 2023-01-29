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
	// ErrNoSlashInEndpoint ...
	ErrNoSlashInEndpoint = errors.New("ErrNoSlashInEndpoint")
	// ErrNotFindEndpoint ...
	ErrNotFindEndpoint = errors.New("ErrNotFindEndpoint")
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

// GetTargetURL ...
func GetTargetURL(originURL string) (url string, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	serverName := ""
	endpoints, ok := cacheURL2OnlineEndPoint[originURL]
	if !ok {
		err = ErrNotFindEndpoint
		return
	}
	endpoint := ""
	for k := range endpoints {
		// for map是乱序的
		endpoint = k
		break
	}
	if endpoint == "" {
		err = ErrNotFindEndpoint
		return
	}
	idx := strings.IndexByte(endpoint, '/')
	if idx == -1 {
		err = ErrNoSlashInEndpoint
		return
	}

	serverName = endpoint[0:idx]
	targetURL := endpoint[idx:]

	serverAddrs, ok := cacheSvcName2OnlineSvcAddr[serverName]
	if !ok {
		err = ErrNotFindServerAddr
		return
	}
	serverAddr := ""
	for k := range serverAddrs {
		// for map是乱序的
		serverAddr = k
		break
	}
	if serverAddr == "" {
		err = ErrNotFindServerAddr
		return
	}

	return serverAddr + targetURL, nil
}

func setEndpoint(url, endpoint string, status bool) error {
	online := "off"
	if status {
		online = "on"
	}
	mydb.Model(modelURL2EndPoint{}).Save(modelURL2EndPoint{
		URL:      url,
		EndPoint: endpoint,
		Status:   online,
	})
	err := mydb.Error
	mydb.Error = nil
	SyncURL2EndPoint()
	return err
}
