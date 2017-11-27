package alipay


//统一收单交易查询接口
//https://docs.open.alipay.com/api_1/alipay.trade.query

type TradeQueryRequest struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 支付宝交易号
	response     *TradeQueryResponse
}

func (this *TradeQueryRequest) Method() string {
	return "alipay.trade.query"
}

func (this *TradeQueryRequest) Params() map[string]string {
	var m = make(map[string]string)
	if this.AppAuthToken != "" {
		m["app_auth_token"] = this.AppAuthToken
	}
	return m
}

func (this *TradeQueryRequest) Name() string {
	return "biz_content"
}

func (this *TradeQueryRequest) JSON() string {
	return marshal(this)
}

func (this *TradeQueryRequest) GetResponse() Response {
	if this.response == nil {
		this.response = &TradeQueryResponse{}
	}
	return this.response
}

type TradeQueryResponse struct {
	TradeQuery struct {
		Code                string           `json:"code"`
		Msg                 string           `json:"msg"`
		SubCode             string           `json:"sub_code"`
		SubMsg              string           `json:"sub_msg"`
		BuyerLogonId        string           `json:"buyer_logon_id"`                // 买家支付宝账号
		BuyerPayAmount      float64          `json:"buyer_pay_amount,string"`       // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string           `json:"buyer_user_id"`                 // 买家在支付宝的用户id
		InvoiceAmount       float64          `json:"invoice_amount,string"`         // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		Openid              string           `json:"open_id"`                       // 买家支付宝用户号，该字段将废弃，不要使用
		OutTradeNo          string           `json:"out_trade_no"`                  // 商家订单号
		PointAmount         float64          `json:"point_amount,string"`           // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount       float64          `json:"receipt_amount,string"`         // 实收金额，单位为元，两位小数
		SendPayDate         string           `json:"send_pay_date"`                 // 本次交易打款给卖家的时间
		TotalAmount         float64          `json:"total_amount,string"`           // 交易的订单金额
		TradeNo             string           `json:"trade_no"`                      // 支付宝交易号
		TradeStatus         string           `json:"trade_status"`                  // 交易状态
		AliPayStoreId       string           `json:"alipay_store_id"`               // 支付宝店铺编号
		StoreId             string           `json:"store_id"`                      // 商户门店编号
		TerminalId          string           `json:"terminal_id"`                   // 商户机具终端编号
		StoreName           string           `json:"store_name"`                    // 请求交易支付中的商户店铺的名称
		DiscountGoodsDetail string           `json:"discount_goods_detail"`         // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string           `json:"industry_sepc_detail"`          // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
		FundBillList        []*FundBill      `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
		VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response,omitempty"`
	ErrorResponse `json:"error_response,omitempty"`
	Sign          string `json:"sign"`
}

type FundBill struct {
	FundChannel string  `json:"fund_channel"`       // 交易使用的资金渠道，详见 支付渠道列表
	Amount      string  `json:"amount"`             // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount,string"` // 渠道实际付款金额
}

type VoucherDetail struct {
	Id                 string `json:"id"`                  // 券id
	Name               string `json:"name"`                // 券名称
	Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
	Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string `json:"memo"`                // 优惠券备注信息
}

func (this *TradeQueryResponse) IsSuccess() bool {
	if this.TradeQuery.Code == "10000" {
		return true
	}
	return false
}
