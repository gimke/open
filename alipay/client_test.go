package alipay

import (
	"testing"
	"fmt"
)

func TestClient(t *testing.T) {
	client := NewClient("https://openapi.alipaydev.com/gateway.do","app_id","your private_key","alipay public_key","RSA2")

	request := &OauthTokenRequest{}
	request.GrantType="authorization_code"
	request.Code="4b203fe6c11548bcabd8da5bb087a83b"

	response := client.Excute(request).(*OauthTokenResponse)

	if response.IsSuccess() {
		fmt.Println(response.Sign)
	}
}