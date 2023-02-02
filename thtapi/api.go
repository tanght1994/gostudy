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
	var user string
	var err error

	endpoint, accessGroup, noPermissionCheck, err := db.GetURLProxy(r.URL.Path)

	// url是否免验证
	if !noPermissionCheck {
		// 验证登录
		user, err = db.ParseToken(r.Header.Get("token"))
		if err != nil {
			w.WriteHeader(403)
			return
		}

		// 验证权限
		havePermission := false
		groups := db.GetUserGroups(user)
		for _, v1 := range groups {
			for _, v2 := range accessGroup {
				if v1 == v2 {
					havePermission = true
				}
			}
		}
		if !havePermission {
			w.WriteHeader(403)
			return
		}
	}

	req, _ := http.NewRequest(r.Method, endpoint, r.Body)
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
