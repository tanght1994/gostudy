package db

import (
	"errors"
	"fmt"
	"net"
	"regexp"
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

// YamlURLProxy ...
type YamlURLProxy struct {
	URLsInfo map[string]struct {
		IPRule            []string        `yaml:"iprule"`
		NoPermissionCheck bool            `yaml:"no_permission_check"`
		AccessGroup       []int           `yaml:"access_group"`
		Endpoint          map[string]bool `yaml:"endpoint"`
		Comment           string          `yaml:"comment"`
	} `yaml:"urls_info"`
	SvcsInfo map[string]map[string]bool `yaml:"svcs_info"`
}

// URLProxy ...
type URLProxy map[string]OneURLProxy

// OneURLProxy ...
type OneURLProxy struct {
	IPRule []struct {
		allow bool
		ipnet *net.IPNet
	}
	NoPermissionCheck bool
	AccessGroup       []int
	Endpoint          []string
}

func str2ProxyConfig(text string) (urlProxy URLProxy, err error) {
	// 解析配置文件
	data := YamlURLProxy{}
	err = yaml.Unmarshal([]byte(text), &data)
	if err != nil {
		return
	}
	urlProxy = make(URLProxy)

	// 整理SvcsInfo
	onlineSvcsInfo := make(map[string]map[string]struct{})
	reg := regexp.MustCompile(`http[s]?://.+?:\d+`)
	for svcname, addrs := range data.SvcsInfo {
		onlineSvcsInfo[svcname] = make(map[string]struct{})
		for addr, online := range addrs {
			// 检查 addr 是否符合 http[s]://xxx.xxx.xxx.xxx:xxx 这种格式
			if !online {
				continue
			}
			if !reg.Match([]byte(addr)) {
				err = errors.New("1")
				return
			}
			onlineSvcsInfo[svcname][addr] = struct{}{}
		}
	}

	// 整理URLsInfo
	for k, v := range data.URLsInfo {
		// 整理IPRule
		IPRule := []struct {
			allow bool
			ipnet *net.IPNet
		}{}
		for _, v1 := range v.IPRule {
			if len(v1) == 0 {
				err = errors.New("1")
				return
			}
			if !(v1[0] == 'a' || v1[0] == 'd') {
				err = errors.New("2")
				return
			}
			allow := v1[0] == 'a'
			ip := string(v1[1:])
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
			for _, v1 := range AccessGroup {
				if x == v1 {
					return true
				}
			}
			return false
		}
		for _, v1 := range v.AccessGroup {
			if !existence(v1) {
				AccessGroup = append(AccessGroup, v1)
			}
		}

		// 检查Endpoint
		Endpoint := make([]string, 0)
		for endpoint, online := range v.Endpoint {
			if !online {
				continue
			}
			// 检查 endpoint 的格式为 svcname + url
			if len(endpoint) < 2 {
				err = errors.New("1")
				return
			}
			idx := strings.Index(endpoint, "/")
			if idx == -1 || idx == 0 {
				err = errors.New("2")
				return
			}

			// 将endpoint中的svcname替换
			svcname := endpoint[:idx]
			url := endpoint[idx:]
			for addr, _ := range onlineSvcsInfo[svcname] {
				Endpoint = append(Endpoint, addr+url)
			}
		}

		urlProxy[k] = OneURLProxy{
			IPRule:            IPRule,
			NoPermissionCheck: v.NoPermissionCheck,
			AccessGroup:       AccessGroup,
			Endpoint:          Endpoint,
		}
	}
	return
}

// SyncAllCache ...
func SyncAllCache() {
	text := `
urls_info:
  a/b/c/d/e:
    iprule: ['a10.251.0.0/16', 'd192.168.7.45/32']
    no_permission_check: false
    access_group: [1, 2, 3]
    endpoint:
      'server1/a/b': on
      'server2/a/b': on
    comment: xxxxxxxxx
  a/a/a/a/a:
    iprule: ['a10.251.0.0/16', 'd192.168.7.45/32']
    no_permission_check: false
    access_group: [4, 5]
    endpoint:
      'server1/a/b': on
      'server2/a/b': on
    comment: xxxxxxxxx
svcs_info:
  'server1':
    'http://192.168.7.45:8000': on
    'http://192.168.7.45:8001': off
  'server2':
    'http://192.168.7.45:8002': on
    'http://192.168.7.45:8003': on
`
	a, e := str2ProxyConfig(text)
	fmt.Println(a)
	fmt.Println(e)
	return
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
