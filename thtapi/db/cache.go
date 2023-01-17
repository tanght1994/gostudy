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

func updateCache() {
	syncURL2EndPoint()
	syncSvcName2SvcAddr()
	syncUserGroup()
	syncGroupURL()
	syncURLWhiteList()
}

func getTableModifiedTag(tableName string) int64 {
	tag := []int64{}
	mydb.Model(modelTableModified{}).Select("tag").Where("name=?", tableName).Find(&tag)
	if len(tag) == 0 {
		return -1
	}
	return tag[0]
}

func syncURL2EndPoint() {
	tag := getTableModifiedTag(modelURL2EndPoint{}.TableName())
	defer func() {
		tagURL2EndPoint = tag
	}()
	if tag == tagURL2EndPoint || tag == -1 {
		return
	}

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

func syncSvcName2SvcAddr() {
	tag := getTableModifiedTag(modelSvcName2SvcAddr{}.TableName())
	defer func() {
		tagURL2EndPoint = tag
	}()
	if tag == tagURL2EndPoint || tag == -1 {
		return
	}

	datas := []modelSvcName2SvcAddr{}
	mydb.Model(modelSvcName2SvcAddr{}).Where("status=?", "on").Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheSvcName2SvcAddr[data.SvcName] = append(cacheSvcName2SvcAddr[data.SvcName], data.SvcAddr)
	}
	common.LogCritical("sync cacheSvcName2SvcAddr ok")
}

func syncUserGroup() {
	tag := getTableModifiedTag(modelUserGroup{}.TableName())
	defer func() {
		tagUserGroup = tag
	}()
	if tag == tagUserGroup || tag == -1 {
		return
	}

	datas := []modelUserGroup{}
	mydb.Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheUserGroup[data.Username] = append(cacheUserGroup[data.Username], data.GroupName)
	}
	common.LogCritical("sync cacheUserGroup ok")
}

func syncGroupURL() {
	tag := getTableModifiedTag(modelGroupURL{}.TableName())
	defer func() {
		tagGroupURL = tag
	}()
	if tag == tagGroupURL || tag == -1 {
		return
	}

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

func syncURLWhiteList() {
	tag := getTableModifiedTag(modelURLWhiteList{}.TableName())
	defer func() {
		tagURLWhiteList = tag
	}()
	if tag == tagURLWhiteList || tag == -1 {
		return
	}

	datas := []modelURLWhiteList{}
	mydb.Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	for _, data := range datas {
		cacheURLWhiteList[data.URL] = struct{}{}
	}
	common.LogCritical("sync cacheURLWhiteList ok")
}
