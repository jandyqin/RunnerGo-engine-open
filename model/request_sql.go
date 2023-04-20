package model

import uuid "github.com/satori/go.uuid"

type SQL struct {
	TargetId      string               `json:"target_id"`
	Uuid          uuid.UUID            `json:"uuid"`
	Name          string               `json:"name"`
	TeamId        string               `json:"team_id"`
	TargetType    string               `json:"target_type"` // api/webSocket/tcp/grpc
	SqlInfo       SqlInfo              `json:"sql_info"`
	Assert        []*AssertionText     `json:"assert"`  // 验证的方法(断言)
	Timeout       int64                `json:"timeout"` // 请求超时时间
	Regex         []*RegularExpression `json:"regex"`   // 正则表达式
	Debug         string               `json:"debug"`   // 是否开启Debug模式
	Configuration *Configuration       `json:"configuration"`
	Variable      *GlobalVariable      `json:"variable"` // 全局变量
}

type SqlInfo struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	Charset  string `json:"charset"`
}
