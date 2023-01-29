package main

import (
	"io/ioutil"
	"net/http"
	"thtapi/db"

	"github.com/gin-gonic/gin"
)

func proxy(ginctx *gin.Context) {
	w := ginctx.Writer
	r := ginctx.Request
	var user, targetURL, originURL string
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

	targetURL, err = db.GetTargetURL(originURL)
	if err == db.ErrNotFindTargetURL {
		w.WriteHeader(404)
		return
	} else if err == db.ErrNotFindServerAddr {
		w.WriteHeader(502)
		return
	}

	req, _ := http.NewRequest(r.Method, targetURL, r.Body)
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

func setEndPoint(ginctx *gin.Context) {
	type Data struct {
		URL      string `binding:"required" json:"url"`
		EndPoint string `binding:"required" json:"endpoint"`
		Status   string `binding:"required" json:"status"`
	}
	data := Data{}
	err := ginctx.BindJSON(&data)
	if err != nil {
		return
	}
	ginctx.Data(200, "text/plain", []byte("haha hehe"))
}
