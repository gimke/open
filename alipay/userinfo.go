package alipay

//https://docs.open.alipay.com/api_9/alipay.system.oauth.token

type UserInfoShareRequest struct {
	Token    string `json:"-"` //
	response *UserInfoShareResponse
}

func (this *UserInfoShareRequest) Method() string {
	return "alipay.user.info.share"
}

func (this *UserInfoShareRequest) Params() map[string]string {
	var m = make(map[string]string)
	if this.Token != "" {
		m["auth_token"] = this.Token
	}
	return m
}

func (this *UserInfoShareRequest) Name() string {
	return ""
}

func (this *UserInfoShareRequest) JSON() string {
	return ""
}

func (this *UserInfoShareRequest) GetResponse() Response {
	if this.response == nil {
		this.response = &UserInfoShareResponse{}
	}
	return this.response
}

type UserInfoShareResponse struct {
	UserInfoShareResponse struct {
		Code     string `json:"code"`
		Msg      string `json:"msg"`
		Avatar   string `json:"avatar"`
		Province string `json:"province"`
		UserId   string `json:"user_id"`   //
		NickName string `json:"nick_name"` //
		Gender   string `json:"gender"`    //
	} `json:"alipay_user_info_share_response,omitempty"`
	ErrorResponse `json:"error_response,omitempty"`
	Sign          string `json:"sign"`
}

func (this *UserInfoShareResponse) IsSuccess() bool {
	if this.UserInfoShareResponse.Code == "10000" {
		return true
	}
	return false
}
