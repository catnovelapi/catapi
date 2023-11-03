package catapi

import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

const deviceToken = "ciweimao_"
const useragent = "Android com.kuangxiangciweimao.novel "

type CiweimaoClient struct {
	Ciweimao *catapi.Ciweimao
}

func NewCiweimaoClient() *CiweimaoClient {
	client := &CiweimaoClient{&catapi.Ciweimao{}}
	client.Ciweimao.Req = &catapi.CiweimaoRequest{
		Debug:         false,
		Version:       "2.9.290",
		BuilderClient: resty.New().SetRetryCount(7).SetBaseURL("https://app.hbooker.com"),
	}
	if client.Ciweimao.Req.Proxy != "" {
		client.Ciweimao.Req.BuilderClient.SetProxy(client.Ciweimao.Req.Proxy)
	}
	client.Ciweimao.Req.BuilderClient.SetFormData(map[string]string{
		"device_token": deviceToken,
		"app_version":  client.Ciweimao.Req.Version,
		"login_token":  client.Ciweimao.Req.LoginToken,
		"account":      client.Ciweimao.Req.Account,
	})

	client.Ciweimao.Req.BuilderClient.SetHeaders(map[string]string{"User-Agent": useragent + client.Ciweimao.Req.Version})
	return client
}

func (ciweimaoClient *CiweimaoClient) SetVersion(version string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Version = version
	return ciweimaoClient
}

func (ciweimaoClient *CiweimaoClient) SetDebug() *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Debug = true
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetDebug(true)
	file, err := os.OpenFile("catapi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open file error !")
	}
	ciweimaoClient.Ciweimao.Req.FileLog = file
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetProxy(proxy string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Proxy = proxy
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetLoginToken(loginToken string) *CiweimaoClient {
	if len(loginToken) != 32 {
		log.Println("loginToken length is not 32")
	} else {
		ciweimaoClient.Ciweimao.Req.LoginToken = loginToken
	}
	return ciweimaoClient

}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func (ciweimaoClient *CiweimaoClient) SetAccount(account string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Account = account
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetAuth(account, loginToken string) *CiweimaoClient {
	return ciweimaoClient.SetAccount(account).SetLoginToken(loginToken)
}
