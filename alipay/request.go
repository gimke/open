package alipay

import (
	"encoding/json"
)

type Request interface {
	// 用于提供访问的 method
	Method() string

	// 返回参数列表
	Params() map[string]string

	// 返回扩展 JSON 参数的字段名称
	Name() string

	// 返回扩展 JSON 参数的字段值
	JSON() string

	GetResponse() Response
}

func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
