package db

import (
	"sync"
	"thtapi/common"
)

var (
	cacheURL2EndPoint    = make(map[string]string)
	cacheSvcName2SvcAddr = make(map[string][]string)
	cacheUserGroup       = make(map[string][]string)
	cacheGroupURL        = make(map[string]map[string]struct{})
	cacheURLWhiteList    = make(map[string]struct{})
	cacheUserToken       = make(map[string]string)

	tagURL2EndPoint    int64
	tagSvcName2SvcAddr int64
	tagUserGroup       int64
	tagGroupURL        int64
	tagURLWhiteList    int64
	tagUserToken       int64

	synclock sync.RWMutex
)

// SyncAllCache ...
func SyncAllCache() {
	SyncURL2EndPoint()
	SyncSvcName2SvcAddr()
	SyncUserGroup()
	SyncGroupURL()
	SyncURLWhiteList()
}

// SyncURL2EndPoint ...
func SyncURL2EndPoint() {
	datas := []modelURL2EndPoint{}
	mydb.Model(modelURL2EndPoint{}).Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	cacheURL2EndPoint = make(map[string]string)
	for _, data := range datas {
		cacheURL2EndPoint[data.URL] = data.EndPoint
	}
	common.LogCritical("sync cacheURL2EndPoint ok")
}

// SyncSvcName2SvcAddr ...
func SyncSvcName2SvcAddr() {
	datas := []modelSvcName2SvcAddr{}
	mydb.Model(modelSvcName2SvcAddr{}).Where("status=?", "on").Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheSvcName2SvcAddr[data.SvcName] = append(cacheSvcName2SvcAddr[data.SvcName], data.SvcAddr)
	}
	common.LogCritical("sync cacheSvcName2SvcAddr ok")
}

// SyncUserGroup ...
func SyncUserGroup() {
	datas := []modelUserGroup{}
	mydb.Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheUserGroup[data.Username] = append(cacheUserGroup[data.Username], data.GroupName)
	}
	common.LogCritical("sync cacheUserGroup ok")
}

// SyncGroupURL ...
func SyncGroupURL() {
	datas := []modelGroupURL{}
	mydb.Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		if _, ok := cacheGroupURL[data.GroupName]; !ok {
			cacheGroupURL[data.GroupName] = make(map[string]struct{})
		}
		cacheGroupURL[data.GroupName][data.URL] = struct{}{}
	}
	common.LogCritical("sync cacheGroupURL ok")
}

// SyncURLWhiteList ...
func SyncURLWhiteList() {
	datas := []modelURLWhiteList{}
	mydb.Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheURLWhiteList[data.URL] = struct{}{}
	}
	common.LogCritical("sync cacheURLWhiteList ok")
}
