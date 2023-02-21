package main

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/sourcegraph/conc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// AAA ...
type AAA struct{}

// LogDebug ...
func (a AAA) LogDebug(msg string) {}

// LogInfo ...
func (a AAA) LogInfo(msg string) {}

// LogWarning ...
func (a AAA) LogWarning(msg string) {}

// LogError ...
func (a AAA) LogError(msg string) {}

// LogCritical ...
func (a AAA) LogCritical(msg string) {}

func hahahahaha() {
	discover, err := NewDiscover("/discover", AAA{})
	if err != nil {
		return
	}
	wg := conc.NewWaitGroup()
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			for j := 0; j < 10000; j++ {
				discover.GetEndpoint("svc1")
			}
		})
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(t0))
}

// IDiscoverLog ...
type IDiscoverLog interface {
	LogDebug(msg string)
	LogInfo(msg string)
	LogWarning(msg string)
	LogError(msg string)
	LogCritical(msg string)
}

// NewDiscover ...
func NewDiscover(prefix string, logger IDiscoverLog) (*Discover, error) {
	discover := &Discover{
		prefix:      prefix,
		svc2eps1:    make(map[string][]EndpointWeight),
		svc2eps2:    make(map[string][]EndpointWeight),
		totalWeight: make(map[string]int),
		lock:        sync.RWMutex{},
		logger:      logger,
	}
	err := discover.watch(context.TODO())
	if err != nil {
		return nil, err
	}
	return discover, nil
}

// Discover ...
type Discover struct {
	prefix      string
	svc2eps1    map[string][]EndpointWeight // 服务名对应的IP地址列表(包含权重为0的IP地址)
	svc2eps2    map[string][]EndpointWeight // 服务名对应的IP地址列表(不包含权重为0的IP地址)
	totalWeight map[string]int
	lock        sync.RWMutex
	logger      IDiscoverLog
}

// EndpointWeight ...
type EndpointWeight struct {
	Endpoint string
	Weight   int
}

// GetEndpoint ...
func (d *Discover) GetEndpoint(server string) string {
	d.lock.RLock()
	defer d.lock.RUnlock()
	eps, ok := d.svc2eps2[server]
	if !ok || len(eps) == 0 {
		d.logger.LogDebug(fmt.Sprintf("GetEndpoint(%s) 1, ok is[%v], len(eps) is [%v]", server, ok, len(eps)))
		return ""
	}
	totalWeight := d.totalWeight[server]
	if totalWeight == 0 {
		d.logger.LogDebug(fmt.Sprintf("GetEndpoint(%s) 2, totalWeight == 0", server))
		return ""
	}
	r := rand.Intn(totalWeight)
	w := 0
	for i := range eps {
		w += eps[i].Weight
		if r < w {
			return eps[i].Endpoint
		}
	}
	d.logger.LogWarning(fmt.Sprintf("GetEndpoint(%s) 3", server))
	return eps[0].Endpoint
}

func (d *Discover) watch(ctx context.Context) error {
	kreg := regexp.MustCompile(fmt.Sprintf("%s/([^/]+)/.+", d.prefix))
	vreg := regexp.MustCompile("([^ ]+) ([0-9]+)")

	timeout := 5 * time.Second
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"www.tanght.xyz:2379"},
		DialTimeout: timeout,
		Context:     ctx,
		DialOptions: []grpc.DialOption{
			grpc.WithTimeout(timeout),
			grpc.WithBlock(),
		},
	})
	if err != nil {
		return err
	}

	getAllKV := func() (r *clientv3.GetResponse, err error) {
		return clientv3.NewKV(cli).Get(ctx, d.prefix, clientv3.WithPrefix())
	}

	r, err := getAllKV()
	if err != nil {
		return err
	}

	dealK := func(k string) (name string, err error) {
		// k格式 prefix/servername/xxxx
		r := kreg.FindAllStringSubmatch(k, 1)
		if r == nil {
			return "", fmt.Errorf("123")
		}
		return r[0][1], nil
	}

	dealV := func(v string) (endPoint string, weight int, err error) {
		// v格式 endpiint weight
		r := vreg.FindAllStringSubmatch(v, 1)
		if r == nil {
			return "", 0, fmt.Errorf("123")
		}
		w, e := strconv.Atoi(r[0][2])
		if e != nil {
			return "", 0, e
		}
		return r[0][1], w, nil
	}

	onGetResponse := func(r *clientv3.GetResponse) {
		svc2eps1 := make(map[string][]EndpointWeight)
		svc2eps2 := make(map[string][]EndpointWeight)
		totalWeight := make(map[string]int)
		for i := range r.Kvs {
			name, err := dealK(string(r.Kvs[i].Key))
			if err != nil {
				d.logger.LogError(fmt.Sprintf("Discover.watch.onGetResponse 1, %s", err.Error()))
				continue
			}
			endpoint, weight, err := dealV(string(r.Kvs[i].Value))
			if err != nil {
				d.logger.LogError(fmt.Sprintf("Discover.watch.onGetResponse 2, %s", err.Error()))
				continue
			}
			svc2eps1[name] = append(svc2eps1[name], EndpointWeight{Endpoint: endpoint, Weight: weight})
			if weight == 0 {
				continue
			}
			svc2eps2[name] = append(svc2eps2[name], EndpointWeight{Endpoint: endpoint, Weight: weight})
			totalWeight[name] += weight
		}
		d.lock.Lock()
		defer d.lock.Unlock()
		d.totalWeight = totalWeight
		d.svc2eps1 = svc2eps1
		d.svc2eps2 = svc2eps2
	}

	onGetResponse(r)

	// onWatchResponse := func(r *clientv3.WatchResponse) {
	// 	es := r.Events
	// 	for i := range es {
	// 		if es[i].IsCreate() {
	// 			name, err := dealK(string(es[i].Kv.Key))
	// 			if err != nil {
	// 				// log error
	// 				continue
	// 			}
	// 			endpoint, weight, err := dealV(string(es[i].Kv.Value))
	// 			if err != nil {
	// 				// log error
	// 				continue
	// 			}
	// 			d.lock.Lock()
	// 			d.svc2eps1[name] = append(d.svc2eps1[name], EndpointWeight{Endpoint: endpoint, Weight: weight})
	// 			if weight != 0 {
	// 				d.svc2eps2[name] = append(d.svc2eps2[name], EndpointWeight{Endpoint: endpoint, Weight: weight})
	// 				d.totalWeight += weight
	// 			}
	// 			d.lock.Unlock()
	// 		} else if es[i].IsModify() {
	// 		} else {
	// 		}
	// 	}
	// }

	ch := cli.Watch(ctx, d.prefix, clientv3.WithPrefix(), clientv3.WithRev(r.Header.Revision+1))
	go func() {
		for {
			select {
			case <-ch:
				// etcd中的数据发生变化, 直接重新全量获取一次, 简单粗暴无BUG
				// onWatchResponse(&r)
				r, err = getAllKV()
				if err != nil {
					d.logger.LogError(fmt.Sprintf("Discover.watch 1, %s", err.Error()))
					continue
				}
				onGetResponse(r)
			case <-ctx.Done():
				break
			}
		}
	}()
	return nil
}

// Show ...
func (d *Discover) Show() {
	d.lock.RLock()
	defer d.lock.RUnlock()
	fmt.Println(time.Now())
	fmt.Println("svc2eps1 ----")
	for k, v := range d.svc2eps1 {
		for i := range v {
			fmt.Printf("%-10s %-5d %s\n", k, v[i].Weight, v[i].Endpoint)
		}
	}
	fmt.Println("svc2eps2 ----")
	for k, v := range d.svc2eps2 {
		for i := range v {
			fmt.Printf("%-10s %-5d %s\n", k, v[i].Weight, v[i].Endpoint)
		}
	}
}
