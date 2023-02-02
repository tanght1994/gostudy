package db

import (
	"errors"
	"net"
	"regexp"
	"strings"
	"sync"
	"thtapi/common"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

var (
	cacheUserGroup   = make(map[string][]uint32)
	cacheUserToken   = make(map[string]string)
	cacheProxyConfig = make(ProxyConfig)

	synclock sync.RWMutex
)

// YamlProxyConfig ...
type YamlProxyConfig struct {
	URLsInfo map[string]struct {
		IPRule            []string        `yaml:"iprule"`
		NoPermissionCheck bool            `yaml:"no_permission_check"`
		AccessGroup       []uint32        `yaml:"access_group"`
		Endpoint          map[string]bool `yaml:"endpoint"`
		Comment           string          `yaml:"comment"`
	} `yaml:"urls_info"`
	SvcsInfo map[string]map[string]bool `yaml:"svcs_info"`
}

// ProxyConfig ...
type ProxyConfig map[string]OneURLProxy

// OneURLProxy ...
type OneURLProxy struct {
	IPRule []struct {
		allow bool
		ipnet *net.IPNet
	}
	NoPermissionCheck bool
	AccessGroup       []uint32
	Endpoint          []string
}

// SyncProxyConfig ...
func SyncProxyConfig(text string) error {
	record := modelProxyConfig{}
	if err := mydb.Model(modelProxyConfig{}).Last(&record).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}
	// 解析配置文件
	data := YamlProxyConfig{}
	err := yaml.Unmarshal([]byte(record.Text), &data)
	if err != nil {
		return err
	}
	config := make(ProxyConfig)

	// 整理onlineSvcsInfo
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
				return err
			}
			onlineSvcsInfo[svcname][addr] = struct{}{}
		}
	}

	// 整理URLsInfo
	for originURL, oneInfo := range data.URLsInfo {
		// 整理IPRule
		IPRule := []struct {
			allow bool
			ipnet *net.IPNet
		}{}
		for _, v1 := range oneInfo.IPRule {
			if len(v1) == 0 {
				err = errors.New("1")
				return err
			}
			if !(v1[0] == 'a' || v1[0] == 'd') {
				err = errors.New("2")
				return err
			}
			allow := v1[0] == 'a'
			cidr := string(v1[1:])
			_, ipnet, e := net.ParseCIDR(cidr)
			if e != nil {
				return e
			}
			IPRule = append(IPRule, struct {
				allow bool
				ipnet *net.IPNet
			}{allow: allow, ipnet: ipnet})
		}

		// 整理AccessGroup
		AccessGroup := removeDuplicate(oneInfo.AccessGroup)

		// 检查Endpoint
		Endpoint := make([]string, 0)
		for endpoint, online := range oneInfo.Endpoint {
			if !online {
				continue
			}
			// 检查 endpoint 的格式为 svcname + url
			if len(endpoint) < 2 {
				err = errors.New("1")
				return err
			}
			idx := strings.Index(endpoint, "/")
			if idx == -1 || idx == 0 {
				err = errors.New("2")
				return err
			}

			// 将endpoint中的svcname替换
			svcname := endpoint[:idx]
			url := endpoint[idx:]
			for addr := range onlineSvcsInfo[svcname] {
				Endpoint = append(Endpoint, addr+url)
			}
		}

		config[originURL] = OneURLProxy{
			IPRule:            IPRule,
			NoPermissionCheck: oneInfo.NoPermissionCheck,
			AccessGroup:       AccessGroup,
			Endpoint:          Endpoint,
		}
	}
	synclock.Lock()
	defer synclock.Unlock()
	cacheProxyConfig = config
	return err
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
		cacheUserGroup[data.Username] = append(cacheUserGroup[data.Username], data.GroupID)
	}
	common.LogCritical("sync cacheUserGroup ok")
}

// 数组去重
func removeDuplicate[T comparable](list []T) []T {
	m := make(map[T]struct{})
	r := make([]T, 0)
	for _, v := range list {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}
