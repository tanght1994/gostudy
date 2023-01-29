package db

import (
	"sync"
	"thtapi/common"
)

var (
	cacheURL2EndPoint    = make(map[string]map[string]bool)
	cacheSvcName2SvcAddr = make(map[string]map[string]bool)
	cacheUserGroup       = make(map[string][]string)
	cacheGroupURL        = make(map[string]map[string]struct{})
	cacheURLWhiteList    = make(map[string]struct{})
	cacheUserToken       = make(map[string]string)

	// 更新 cacheURL2EndPoint 时顺便更新 cacheURL2OnlineEndPoint
	cacheURL2OnlineEndPoint = make(map[string]map[string]struct{})
	// 更新 cacheSvcName2SvcAddr 时顺便更新 cacheSvcName2OnlineSvcAddr
	cacheSvcName2OnlineSvcAddr = make(map[string]map[string]struct{})

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
	err := mydb.Model(modelURL2EndPoint{}).Find(&datas).Error
	if err != nil {
		common.LogError("SyncURL2EndPoint error, " + err.Error())
		return
	}
	synclock.Lock()
	defer synclock.Unlock()
	cacheURL2EndPoint = make(map[string]map[string]bool)
	cacheURL2OnlineEndPoint = make(map[string]map[string]struct{})
	for _, data := range datas {
		online := data.Status == "on"
		cacheURL2EndPoint[data.URL][data.EndPoint] = online
		if online {
			cacheURL2OnlineEndPoint[data.URL][data.EndPoint] = struct{}{}
		}
	}
	common.LogCritical("sync cacheURL2EndPoint ok")
}

// SyncSvcName2SvcAddr ...
func SyncSvcName2SvcAddr() {
	datas := []modelSvcName2SvcAddr{}
	err := mydb.Model(modelSvcName2SvcAddr{}).Find(&datas).Error
	if err != nil {
		common.LogError("SyncSvcName2SvcAddr error, " + err.Error())
		return
	}
	synclock.Lock()
	defer synclock.Unlock()
	cacheSvcName2SvcAddr = make(map[string]map[string]bool)
	cacheSvcName2OnlineSvcAddr = make(map[string]map[string]struct{})
	for _, data := range datas {
		online := data.Status == "on"
		cacheSvcName2SvcAddr[data.SvcName][data.SvcAddr] = online
		if online {
			cacheSvcName2OnlineSvcAddr[data.SvcName][data.SvcAddr] = struct{}{}
		}
	}
	common.LogCritical("sync cacheSvcName2SvcAddr ok")
}

// SyncUserGroup ...
func SyncUserGroup() {
	datas := []modelUserGroup{}
	err := mydb.Find(&datas).Error
	if err != nil {
		common.LogError("SyncUserGroup error, " + err.Error())
		return
	}
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
	err := mydb.Find(&datas).Error
	if err != nil {
		common.LogError("SyncGroupURL error, " + err.Error())
		return
	}
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
	err := mydb.Find(&datas).Error
	if err != nil {
		common.LogError("SyncURLWhiteList error, " + err.Error())
		return
	}
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheURLWhiteList[data.URL] = struct{}{}
	}
	common.LogCritical("sync cacheURLWhiteList ok")
}
