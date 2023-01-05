package db

import (
	"errors"
	"strings"

	"gorm.io/gorm"
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
	_, ok := urlWhiteList[url]
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
	groups, ok := userGroup[username]
	if !ok {
		return false
	}
	for _, group := range groups {
		if permission, ok1 := groupPermission[group]; ok1 {
			if _, ok2 := permission[url]; ok2 {
				return true
			}
		}
	}
	return false
}

// GetTargetURL ...
func GetTargetURL(originURL string) (serverAddr []string, targetURL string, err error) {
	synclock.RLock()
	defer synclock.RUnlock()
	serverName := ""
	s, ok := url2url[originURL]
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
	serverAddr, ok = serverName2Addr[serverName]
	if !ok {
		err = ErrNotFindServerAddr
		return
	}
	return
}

// tableTagInc table_modified表tag字段+1
func tableTagInc(tableName string) {
	mydb.Model(TableModified{}).Where("name=?", tableName).Update("tag", gorm.Expr("tag+?", 1))
}
