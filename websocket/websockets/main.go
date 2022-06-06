package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handlews)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	must(err, "http.ListenAndServe error")
}

var Upgrade = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var SessionManger = &TSessionManger{}

type TSessionManger struct {
	sid2s      sync.Map // session id to session [uint64: *TSession]
	active_cnt int64
	total_cnt  uint64
	close_cnt  uint64
}

func (sm *TSessionManger) CreateSession(wc *websocket.Conn) *TSession {
	sid := atomic.AddUint64(&sm.total_cnt, 1)
	s := &TSession{sid: sid, wc: wc, once: sync.Once{}, sm: sm}
	sm.sid2s.Store(sid, s)
	fmt.Printf("session[%d] create successful\n", s.sid)
	return s
}

func (sm *TSessionManger) GetSession(sid uint64) (*TSession, bool) {
	s, ok := sm.sid2s.Load(sid)
	return s.(*TSession), ok
}

type TSession struct {
	sid  uint64          // session id
	wc   *websocket.Conn // websocket conn
	once sync.Once
	sm   *TSessionManger // session manger
}

func (s *TSession) Close() {
	s.once.Do(func() {
		var err error
		sid := s.sid
		defer func() {
			if err != nil {
				fmt.Printf("session[%d] close error, %v\n", sid, err)
			} else {
				fmt.Printf("session[%d] close successful\n", sid)
			}
		}()
		err = s.wc.Close()
		s.sm.sid2s.Delete(sid)
		atomic.AddUint64(&s.sm.close_cnt, 1)
		atomic.AddInt64(&s.sm.active_cnt, -1)
	})
}

func (s *TSession) WriteJson(data interface{}) error {
	err := s.wc.WriteJSON(data)
	if err == nil {
		return nil
	}
	fmt.Printf("session[%d] WriteJson error, %v", s.sid, err)
	s.Close()
	return err
}

func (s *TSession) Read() ([]byte, error) {
	_, data, err := s.wc.ReadMessage()
	if err == nil {
		return data, nil
	}
	fmt.Printf("session[%d] read error, %v", s.sid, err)
	s.Close()
	return nil, err
}

func handlews(rsp http.ResponseWriter, req *http.Request) {
	wc, err := Upgrade.Upgrade(rsp, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建session
	s := SessionManger.CreateSession(wc)
	defer s.Close()

	// 将 session id 发送给客户端
	if err := s.WriteJson(NewSTCMsg(STCMD_YourInfo{Uid: s.sid})); err != nil {
		return
	}

	// 循环读取客户端发送的数据
	for {
		msg := &CTSMsg{}
		err = wc.ReadJSON(msg)
		if err != nil {
			break
		}
		handler, ok := event_handle[msg.Ty]
		if !ok {
			break
		}
		handler(s, msg.Data)
	}
}

func must(e error, msg string) {
	if e != nil {
		fmt.Println(msg, e)
		panic(e)
	}
}

type event_func func(*TSession, string) error

var event_handle = map[int]event_func{
	CTSMT_PTPChat: func(self *TSession, data string) error {
		ptpchat := &CTSMD_PTPChat{}
		if err := json.Unmarshal([]byte(data), ptpchat); err != nil {
			return err
		}
		tid := uint64(ptpchat.To)
		if self.sid == tid {
			return nil
		}
		str := ptpchat.Msg
		target, ok := SessionManger.GetSession(tid)
		if !ok {
			return nil
		}
		return target.WriteJson(NewSTCMsg(STCMD_PTPChat{From: self.sid, To: target.sid, Msg: str}))
	},
}
