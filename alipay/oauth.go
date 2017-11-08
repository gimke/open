package alipay

//https://docs.open.alipay.com/api_9/alipay.system.oauth.token

type OauthTokenRequest struct {
	GrantType		string  `json:"-"` 	// 必须 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code			string  `json:"-"` 	// 可选 授权码，用户对应用授权后得到。
	RefreshToken	string  `json:"-"`  // 可选 刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
	response		*OauthTokenResponse
}

func (r *OauthTokenRequest) Method() string {
	return "alipay.system.oauth.token"
}

func (r *OauthTokenRequest) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = r.GrantType
	if r.Code != "" {
		m["code"] = r.Code
	}
	if r.RefreshToken != "" {
		m["refresh_token"] = r.RefreshToken
	}
	return m
}

func (this *OauthTokenRequest) Name() string {
	return ""
}

func (this *OauthTokenRequest) JSON() string {
	return ""
}

func (this *OauthTokenRequest) GetResponse() Response {
	if this.response == nil {
		this.response = &OauthTokenResponse{}
	}
	return this.response
}


type OauthTokenResponse struct {
	OauthTokenResponse struct {
		Code			string	`json:"code"`
		Msg				string 	`json:"msg"`
		SubCode			string 	`json:"sub_code"`
		SubMsg			string 	`json:"sub_msg"`
		UserId			string 	`json:"user_id"`     	// 支付宝用户的唯一userId
		AccessToken		string 	`json:"access_token"` 	// 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn		string 	`json:"expires_in"`   	// 访问令牌的有效时间，单位是秒。
		RefreshToken	string 	`json:"refresh_token"`	// 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn		string 	`json:"re_expires_in"`	// 刷新令牌的有效时间，单位是秒。

	} `json:"alipay_system_oauth_token_response"`
	Sign	string `json:"sign"`
}

func (this *OauthTokenResponse) IsSuccess() bool {
	if this.OauthTokenResponse.Code == "10000" {
		return true
	}
	return false
}
