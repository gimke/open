package alipay

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	gateway	       	string
	appId           string
	privateKey      string
	aliPayPublicKey string
	signType        string
	AppAuthToken 	string
}

//创建client
func NewClient(gateway, appId, privateKey, aliPayPublicKey, signType string) *Client {
	return &Client{
		gateway:gateway,
		appId:appId,
		privateKey:privateKey,
		aliPayPublicKey:aliPayPublicKey,
		signType:signType,
	}
}

//执行
func (this *Client) Excute(request Request) Response {

	data := `{
		"alipay_system_oauth_token_response": {
			"code": "10000",
			"user_id": "2088102150477652",
			"access_token": "20120823ac6ffaa4d2d84e7384bf983531473993",
			"expires_in": "3600",
			"refresh_token": "20120823ac6ffdsdf2d84e7384bf983531473993",
			"re_expires_in": "3600"
		},
		"sign": "ERITJKEIJKJHKKKKKKKHJEREEEEEEEEEEE"
	}`
	response := request.GetResponse()
	err := json.Unmarshal([]byte(data), response)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return response
}
