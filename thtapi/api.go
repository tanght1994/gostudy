package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"thtapi/db"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	var user, targetURL, originURL string
	var serverAddrs []string
	var err error
	originURL = r.URL.Path

	// url是否免验证
	if !db.InWhitelist(originURL) {
		// 验证登录
		user, err = db.ParseToken(r.Header.Get("token"))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		// 验证权限
		if !db.HavePermission(user, originURL) {
			w.WriteHeader(400)
			return
		}
	}

	serverAddrs, targetURL, err = db.GetEndPoint(originURL)
	if err == db.ErrNotFindTargetURL {
		w.WriteHeader(404)
		return
	} else if err == db.ErrNotFindServerAddr {
		w.WriteHeader(500)
		return
	}

	serverAddr := serverAddrs[rand.Intn(len(serverAddrs))]
	url := "http://" + serverAddr + targetURL
	req, _ := http.NewRequest(r.Method, url, r.Body)
	req.URL.RawQuery = r.URL.RawQuery
	req.Header = r.Header.Clone()
	req.Header.Del("Connection")
	res, err := Requester.Do(req)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer res.Body.Close()
	header := w.Header()
	for k, v := range res.Header {
		for _, x := range v {
			header.Add(k, x)
		}
	}
	header.Del("Connection")
	data, _ := ioutil.ReadAll(res.Body)
	w.Write(data)
}

func setEndPoint(w http.ResponseWriter, r *http.Request) {}
