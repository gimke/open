package alipay

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gimke/open/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	RSA  = "RSA"
	RSA2 = "RSA2"
)
const (
	RESPONSE_SUFFIX = "_response"
	ERROR_RESPONSE  = "error_response"
	SIGN_NODE_NAME  = "sign"
)

type Client struct {
	gateway         string
	appId           string
	privateKey      string
	aliPayPublicKey string
	signType        string
	AppAuthToken    string
	client          *http.Client
}

//创建client
func NewClient(gateway, appId, privateKey, aliPayPublicKey, signType string) *Client {
	return &Client{
		gateway:         gateway,
		appId:           appId,
		privateKey:      privateKey,
		aliPayPublicKey: aliPayPublicKey,
		signType:        signType,
	}
}

//执行
func (this *Client) Excute(request Request) (response Response, err error) {
	response = request.GetResponse()
	buf, err := this.MakeBuffer(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", this.gateway, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	if this.client == nil {
		this.client = http.DefaultClient
	}
	resp, err := this.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(this.aliPayPublicKey) > 0 {
		var dataStr = string(data)

		var rootNodeName = strings.Replace(request.Method(), ".", "_", -1) + RESPONSE_SUFFIX

		var rootIndex = strings.LastIndex(dataStr, rootNodeName)
		var errorIndex = strings.LastIndex(dataStr, ERROR_RESPONSE)

		var content string
		var sign string

		if rootIndex > 0 {
			content, sign = parserJSONSource(dataStr, rootNodeName, rootIndex)
		} else if errorIndex > 0 {
			content, sign = parserJSONSource(dataStr, ERROR_RESPONSE, errorIndex)
		} else {
			return nil, errors.New("error format")
		}

		if ok, err := verifyResponseData([]byte(content), this.signType, sign, this.aliPayPublicKey); ok == false {
			return nil, err
		}
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (this *Client) MakeBuffer(request Request) (buf io.Reader, err error) {
	var p = url.Values{}
	p.Add("app_id", this.appId)
	p.Add("method", request.Method())
	p.Add("format", "JSON")
	p.Add("charset", "utf-8")
	p.Add("sign_type", this.signType)
	p.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Add("version", "1.0")

	if len(request.Name()) > 0 {
		p.Add(request.Name(), request.JSON())
	}

	var ps = request.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
		}
	}

	var keys = make([]string, 0, 0)
	for key, _ := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var sign string
	if this.signType == RSA {
		sign, err = signRSA(keys, p, []byte(this.privateKey))
	} else {
		sign, err = signRSA2(keys, p, []byte(this.privateKey))
	}
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)

	buf = strings.NewReader(p.Encode())

	return buf, nil
}

func parserJSONSource(rawData string, nodeName string, nodeIndex int) (content string, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\""+SIGN_NODE_NAME+"\"")
	var dataEndIndex = signIndex - 1

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	var signStartIndex = signIndex + len(SIGN_NODE_NAME) + 4
	sign = rawData[signStartIndex:]
	var signEndIndex = strings.LastIndex(sign, "\"}")
	sign = sign[:signEndIndex]

	return content, sign
}

func verifyResponseData(data []byte, signType, sign string, key string) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if signType == RSA {
		err = encoding.VerifyPKCS1v15(data, signBytes, []byte(key), crypto.SHA1)
	} else {
		err = encoding.VerifyPKCS1v15(data, signBytes, []byte(key), crypto.SHA256)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func signRSA2(keys []string, param url.Values, privateKey []byte) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var src = strings.Join(pList, "&")
	sig, err := encoding.SignPKCS1v15([]byte(src), privateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func signRSA(keys []string, param url.Values, privateKey []byte) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var src = strings.Join(pList, "&")
	sig, err := encoding.SignPKCS1v15([]byte(src), privateKey, crypto.SHA1)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}
