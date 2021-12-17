package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var stat_data = []string{
	`{"method": "app_start", "data": {"game_id": 1, "uid": 222222222, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "isfirstin": 1}}`,
	`{"method": "app_end", "data": {"game_id": 1, "uid": 333333333, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "length": 1}}`,
	`{"method": "purchase_success", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "currency": "aaa", "price": 1.1111, "vip": 1, "rp": 1, "chips": 1, "eca_result_id": 1, "myid": 1, "isfirstin": 1}}`,
	`{"method": "purchase_tap", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "myid": 1, "currency": "aaa", "price": 1}}`,
	`{"method": "register", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "myid": 1, "type": "aaa"}}`,
	`{"method": "login", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "type": "aaa", "result": 1, "level": 1, "days": 1, "vip": 1, "chips": 1, "isfirstin": 1}}`,
	`{"method": "game_enter", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "way": "aaa", "game_name": "aaa", "isfirstin": 1, "chips": 1, "level": 1, "vip": 1}}`,
	`{"method": "game_quit", "data": {"game_id": 1, "uid": 222222222, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "way": "aaa", "length": 1, "game_name": "aaa"}}`,
	`{"method": "emoji", "data": {"game_id": 1, "uid": 333333333, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "game_name": "aaa", "type": "aaa", "myid": 1}}`,
	`{"method": "chips_change", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "stats": "aaa", "pid": 1, "time": 1, "reason": "aaa", "debit_amount": 1, "type": "aaa", "debit_balance": 1}}`,
	`{"method": "rp_change", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "stats": "aaa", "pid": 1, "time": 1, "reason": "aaa", "debit_amount": 1, "type": "aaa", "debit_balance": 1}}`,
	`{"method": "match_register", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "game_name": "aaa", "cost": 1, "pid": 1, "result": 1, "reason": "aaa"}}`,
	`{"method": "match_enter", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "game_name": "aaa", "cost": 1, "pid": 1, "time": 1}}`,
	`{"method": "match_result", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "game_name": "aaa", "cost": 1, "pid": 1, "rank": 1, "rewardchips": 1, "rewardrp": 1, "length": 1}}`,
	`{"method": "rate", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "stats": "aaa", "reason": "aaa", "pid": 1, "chips": 1}}`,
	`{"method": "transfer", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "send_id": "aaa", "recive_id": "aaa", "amount": 1, "rate": 1}}`,
	`{"method": "task_complete", "data": {"game_id": 1, "uid": 1111111, "ctime": 111, "country_id": 1, "plat_id": 1, "channel_id": 1, "tostype": "ios11", "devicetype": "asd", "ip": "192.168.7.45", "imei": "hfghdghduj", "version": "1.1.1", "project": "sfgbhdf", "event_name": "event_name", "event_time": 1634655700, "game_name": "aaa", "pid": 1, "type": "aaa", "rewardrp": 1, "model": "aaa"}}`,
}

func stat_main() {
	wg := sync.WaitGroup{}
	for i := 0; i < len(stat_data); i++ {
		wg.Add(1)
		go stat_work(&wg, i)
	}
	wg.Wait()
	fmt.Println("MAIN DONE")
}

func stat_work(wg *sync.WaitGroup, idx int) {
	defer wg.Done()
	url := "http://192.168.7.45:9090/v1/0/msg_data/"
	c := http.Client{Timeout: 5 * time.Second}
	d := stat_data[idx]
	for i := 0; i < 10000; i++ {
		_, err := c.Post(url, "application/json", bytes.NewBuffer([]byte(d)))
		if err != nil {
			fmt.Println(err)
		}
		if (i % 100) == 0 {
			fmt.Println(idx, " --- ", i)
		}
	}
	fmt.Println(idx, " DONE")
}
