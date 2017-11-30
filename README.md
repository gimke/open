# 开放平台

## alipay

### OauthTokenRequest
```go
client := alipay.NewClient(config.Gateway, config.AppId, config.PrivateKey, config.AliPayPublicKey, alipay.RSA2)

request := &alipay.OauthTokenRequest{
    GrantType: "authorization_code",
    Code:      authCode,
}

responseBase, err := client.Excute(request)
if err != nil {
    return
}
response := responseBase.(*alipay.OauthTokenResponse)

if response.IsSuccess() {
    //success
} else {
    //error
}
```

### TradeCreateRequest
```go
client := alipay.NewClient(config.Gateway, config.AppId, config.PrivateKey, config.AliPayPublicKey, alipay.RSA2)

var goodDetail = []*alipay.GoodsDetailItem{}

for _, l := range list {
    goodDetail = append(goodDetail, &alipay.GoodsDetailItem{
        GoodsId:   l.Id.Hex(),
        GoodsName: l.Name,
        Quantity:  strconv.Itoa(l.Count),
        Price:     strconv.FormatFloat(l.Price, 'f', 2, 64),
    })
}
request := &alipay.TradeCreateRequest{
    Subject:        subject,
    OutTradeNo:     outtradeno.Hex(),
    TotalAmount:    amount,
    BuyerId:        user.UserId,
    GoodsDetail:    goodDetail,
    TimeoutExpress: "6h",
}

responseBase, err := client.Excute(request)
if err != nil {
    c.JSON(510, cart.H{"error": err.Error()})
}
response := responseBase.(*alipay.TradeCreateResponse)
if response.IsSuccess() {
    //创建订单
    order := db.NewOrder()
    order.Id = outtradeno
    order.UserId = user.UserId
    order.TradeNo = response.TradeCreate.TradeNo
    order.Source = "alipay"
    order.Amount = amount
    order.StoreId = bson.ObjectIdHex(storeId)
    order.Status = 100
    order.List = list
    success, err := order.Insert()
    if success {
        user.Outtradeno = outtradeno.Hex()
        user.Set(c)
        c.JSON(200, cart.H{
            "tradeno":    response.TradeCreate.TradeNo,
            "outtradeno": outtradeno.Hex(),
        })
    } else {
        c.JSON(510, cart.H{"error": err.Error()})
    }

} else {
    c.JSON(510, cart.H{"error": response.TradeCreate.SubMsg})
}
```
## wechat
