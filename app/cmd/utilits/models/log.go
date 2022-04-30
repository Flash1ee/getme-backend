package models

import (
	"net/url"
	"time"
)

//go:generate easyjson log.go

//easyjson:json
type Log struct {
	Level    string    `json:"level,omitempty"`
	Method   string    `json:"method,omitempty"`
	Msg      string    `json:"msg,omitempty"`
	Adr      string    `json:"remote_addr,omitempty"`
	Url      url.URL   `json:"urls,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	WorkTime int64     `json:"work_time,omitempty"`
	ReqID    string    `json:"req_id,omitempty"`
}
