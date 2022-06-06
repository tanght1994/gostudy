package main

import (
	"encoding/json"
)

const (
	STCMT_Err = iota
	STCMT_YourInfo
	STCMT_PTPChat
)

type STCMsg struct {
	Ty   int    `json:"ty"`
	Data string `json:"data"`
}

type STCMD_YourInfo struct {
	Uid uint64 `json:"uid"`
}

type STCMD_PTPChat struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
	Msg  string `json:"msg"`
}

func NewSTCMsg(data interface{}) STCMsg {
	ty := STCMT_Err
	switch data.(type) {
	case STCMD_YourInfo, *STCMD_YourInfo:
		ty = STCMT_YourInfo
	case STCMD_PTPChat, *STCMD_PTPChat:
		ty = STCMT_PTPChat
	default:
		panic("111")
	}
	bys, err := json.Marshal(data)
	must(err, "json.Marshal(data)")
	return STCMsg{Ty: ty, Data: string(bys)}
}

const (
	CTSMT_Err = iota
	CTSMT_PTPChat
)

type CTSMsg struct {
	Ty   int    `json:"ty"`
	Data string `json:"data"`
}

type CTSMD_PTPChat struct {
	To  uint64 `json:"to"`
	Msg string `json:"msg"`
}

func NewCTSMsg(data interface{}) CTSMsg {
	ty := CTSMT_Err
	switch data.(type) {
	case CTSMD_PTPChat, *CTSMD_PTPChat:
		ty = CTSMT_PTPChat
	default:
		panic("111")
	}
	bys, err := json.Marshal(data)
	must(err, "json.Marshal(data)")
	return CTSMsg{Ty: ty, Data: string(bys)}
}
