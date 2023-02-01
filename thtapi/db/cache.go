package db

import (
	"errors"
	"net"
	"strings"
	"sync"
	"thtapi/common"

	"gopkg.in/yaml.v3"
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

// ProxyConfig ...
type ProxyConfig struct {
	URLsInfo map[string]OneURLInfo
	SvcsInfo map[string]OneSvcInfo
}

// OneURLInfo ...
type OneURLInfo struct {
	IPRule []struct {
		allow bool
		ipnet *net.IPNet
	}
	NoPermissionCheck bool
	AccessGroup       []int
	Endpoint          map[string]bool
	Comment           string
}

// OneSvcInfo ...
type OneSvcInfo map[string]bool

func str2ProxyConfig(text string) (proxyConfig ProxyConfig, err error) {
	// 解析配置文件
	type AAA struct {
		URLsInfo map[string]struct {
			IPRule            []string        `yaml:"iprule"`
			NoPermissionCheck bool            `yaml:"no_permission_check"`
			AccessGroup       []int           `yaml:"access_group"`
			Endpoint          map[string]bool `yaml:"endpoint"`
			Comment           string          `yaml:"comment"`
		} `yaml:"urls_info"`
		SvcsInfo map[string]OneSvcInfo `yaml:"svcs_info"`
	}
	data := AAA{}
	err = yaml.Unmarshal([]byte(text), &data)
	if err != nil {
		return
	}
	proxyConfig.SvcsInfo = make(map[string]OneSvcInfo)
	proxyConfig.URLsInfo = make(map[string]OneURLInfo)

	// 整理svcs_info
	for k, v := range data.SvcsInfo {
		for k1 := range v {
			// 检查是否符合 servername + url 这种格式
			if len(k1) < 2 {
				err = errors.New("1")
				return
			}
			if idx := strings.Index(k1, "/"); (idx == -1) || (idx == 0) {
				err = errors.New("1")
				return
			}
		}
		proxyConfig.SvcsInfo[k] = v
	}

	// 整理urls_info

	for k, v := range data.URLsInfo {
		one := OneURLInfo{}

		// 整理IPRule
		IPRule := []struct {
			allow bool
			ipnet *net.IPNet
		}{}
		for _, v := range v.IPRule {
			if len(v) == 0 {
				err = errors.New("1")
				return
			}
			if !(v[0] == 'a' || v[0] == 'd') {
				err = errors.New("2")
				return
			}
			allow := v[0] == 'a'
			ip := string(v[1:])
			_, ipnet, e := net.ParseCIDR(ip)
			if e != nil {
				err = e
				return
			}
			IPRule = append(IPRule, struct {
				allow bool
				ipnet *net.IPNet
			}{allow: allow, ipnet: ipnet})
		}

		// 整理AccessGroup
		AccessGroup := []int{}
		existence := func(x int) bool {
			for _, v := range one.AccessGroup {
				if x == v {
					return true
				}
			}
			return false
		}
		for _, v := range v.AccessGroup {
			if !existence(v) {
				AccessGroup = append(AccessGroup, v)
			}
		}

		// 整理Endpoint
		Endpoint := []struct {
			URL    string
			Online bool
		}{}
		for _, v := range v.Endpoint {
			Endpoint = append(Endpoint, struct {
				URL    string
				Online bool
			}{URL: "", Online: true})
		}

		allURLInfo[k] = one
	}
	return
}

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
	if err := mydb.Model(modelURL2EndPoint{}).Find(&datas).Error; err != nil {
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
	if err := mydb.Model(modelSvcName2SvcAddr{}).Find(&datas).Error; err != nil {
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
	if err := mydb.Find(&datas).Error; err != nil {
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
	if err := mydb.Find(&datas).Error; err != nil {
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
	if err := mydb.Find(&datas).Error; err != nil {
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
