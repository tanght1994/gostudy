package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"tanght/endpoint"
)

func HelloDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	name := request.URL.Query().Get("name")
	if name == "" {
		return nil, errors.New("无效参数")
	}
	return endpoint.HelloRequest{Name: name}, nil
}

func HelloEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func ByeDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	name := request.URL.Query().Get("name")
	if name == "" {
		return nil, errors.New("无效参数")
	}
	return endpoint.ByeRequest{Name: name}, nil
}

func ByeEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
