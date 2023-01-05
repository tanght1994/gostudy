package db

import "sync"

var (
	url2url         map[string]string
	serverName2Addr map[string][]string
	userGroup       map[string][]string
	groupPermission map[string]map[string]struct{}
	userToken       map[string]string
	urlWhiteList    map[string]struct{}

	tagURL2url         int64
	tagServerName2Addr int64
	tagUserGroup       int64
	tagGroupPermission int64
	tagUserToken       int64
	tagURLWhiteList    int64

	synclock sync.RWMutex
)

func syncdb() {
	syncurl2url()
	syncServerName2Addr()
	syncUserGroup()
	syncGroupPermission()
	syncurlWhiteList()
}

func getTableModifiedTag(tableName string) int64 {
	tag := []int64{}
	mydb.Table(tableName).Select("tag").Where("name=?", tableName).Find(&tag)
	if len(tag) == 0 {
		return -1
	}
	return tag[0]
}

func syncurl2url() {
	tag := getTableModifiedTag(URL2URL{}.TableName())
	defer func() {
		tagURL2url = tag
	}()
	if tag == tagURL2url || tag == -1 {
		return
	}
	// 同步
	datas := []URL2URL{}
	mydb.Model(URL2URL{}).Find(&datas)
	synclock.Lock()
	defer synclock.Unlock()
	url2url = make(map[string]string)
	for _, data := range datas {
		url2url[data.OriginURL] = data.TargetURL
	}
}

func syncServerName2Addr() {
	tag := getTableModifiedTag(ServerAddr{}.TableName())
	defer func() {
		tagURL2url = tag
	}()
	if tag == tagURL2url || tag == -1 {
		return
	}
}

func syncUserGroup() {

}

func syncGroupPermission() {

}

func syncurlWhiteList() {

}
