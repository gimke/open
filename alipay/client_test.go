package alipay

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA14DqJoNbP6yk1rrof4N40xNwwW/mUqP/93OWCToyn2aGKDMWiSvUQRf5Hk7AkY5xqJYBTirPIY6pN4KCvuMJp2cpEYo1Ubd+jgTbUv83GPzyXsbOPl8gOPNa5+IbfqJpihVMe5z7VcOqxgyW8/yuY4H6BoAFq/c77CLKdgh+t1I0fdI+jevNfren1YApcv9y5Igs6sr7/Yq5WXrPOyN5fjx3bT+6TRhJNj9RotuYN4lnGTiineNVmy67WejnI3k48fJEZ75BT4MZ8bqlzLp//RI8pfGmTit/4gsGFZI5DHW5WtsXtDGOecmWsu1NDr8qc/OXYIdbQdOCwtIB083/7wIDAQABAoIBAA6LRcI5pUPj2/AeByjr75CXREnZynqTVOgXZe3Tfq0hVzaJVCEH0zhdbNOtzvND4MnW7dmfrAEAxszXQwms+u7QWIY1hKmyL5lYHJE6ZjaKg7T/x9WPx/Xv/pedu+tM/MBz9Yh2LMQ6A2GAYgOvbvmKQRyFMVzMv7+NDYrvwdWPWAVSbDGE/SRS9f/Er2UVd4OFiogjkI37EQAEOZLgy31K1NXiJq5/aBkLh1WKejA+G+BXp7QKSNrHwZRSPnJelPHIJoZGui884PInzVDKez7VXRH3cEnA7Zwyo1rGqtm6DtaBG4mgpFMvcKH5K/pPCCmntO/bjxG3ILytNGkllSECgYEA62km8s2MTUzZTybGoC7tvs85rCi/WLsJhn1xnnmNZqVuas8nFgHVKrqjuzEYPBjqhh43rHqh6wycB1QJEp4Vmw1hgW6Fzl+GpPLk4dzWYsFJF19NDdvASHwvu7LyOTs/zUFN9+FfT8o7UOzb+Gffm4H5eXWFZ1XydzVmBd9yXv8CgYEA6loKa2ouJjt4dc+79SqNaiwMzpD1UGcqCJpUdQw6qTZRt5wDbN2xDiJ+zLsqOhpSAh6CyDL/2Fmi1KmRL4eWPjKF+nR9H/+cIqGbKg9F/GXUJfOOh1yHJqnZsiOJCvcRqDuWYe/dw/3hSC5C/nEaViA66EoERbBeJ2dXM1MbTxECgYEAmHiHqHUnLR3cFd7ogPFEPPScxvuxSzgBKGFxSJIz2krFpFo9V4yiU0WFLIXUy3/bzjgeGRFodAO7vydXpP5Mwhs9jwZVld/bJlTHl95f4KCNxZyNHK+673e3tttk9VqBrWBhrm4DPHugRX7TziUA+AiL23YZjF9nZnxoct9RhWMCgYEAoHlpuyYpVdrRYPMQBRDPZ02ks2qF4TnTibKMdN4b6TUd/fniSpEAJeqvI2hiwQi28WaNLaPml/LBUpiOp5pT4mFcZyWPbPLuqrQ4+TMePHhKLna6Oay9i1cxkA9PT2fh+m5bStMi64uU0YWEMJGodCN70wakKEheIONdzfJxt9ECgYB5IkDtQvRblmbxa8xpJuVoOxTYG273tpCDA32EtlV/OKokMLIr0sUdfZMKrBhO/ZPCBG9mKck8FjR8vENDayySNPiODUdltU/I09MYHYxtp3nLxCKEqnNsM814UyL3bI6Vj0aaCSKx4l/gOKCqtWn11NZXbmUkNQi7v96XPbaRGg==\n-----END RSA PRIVATE KEY-----"
	client := NewClient("https://openapi.alipaydev.com/gateway.do", "2016090900468760", privateKey, "alipay public_key", RSA2)

	request := &OauthTokenRequest{}
	request.GrantType = "authorization_code"
	request.Code = "4b203fe6c11548bcabd8da5bb087a83b"

	_, err := client.Excute(request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response := request.GetResponse().(*OauthTokenResponse)

		fmt.Println(response)
		if response.IsSuccess() {
			fmt.Println(response.Sign)
		}
	}
}
