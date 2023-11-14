package catapi

import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

type CiweimaoClient struct {
	Ciweimao *catapi.Ciweimao
}

func NewCiweimaoClient() *CiweimaoClient {
	client := &CiweimaoClient{
		Ciweimao: &catapi.Ciweimao{
			Req: &catapi.CiweimaoRequest{Debug: false, BuilderClient: resty.New()},
		},
	}
	client.SetRetryCount(7)
	client.SetBaseURL("https://app.hbooker.com").SetUserAgent("Android")
	var versionNumber string
	if version, err := client.Ciweimao.GetVersionApi(); err != nil {
		versionNumber = "2.9.290"
	} else {
		versionNumber = version
	}
	client.SetVersion(versionNumber)
	client.SetUserAgent("Android com.kuangxiangciweimao.novel")
	return client
}

func (ciweimaoClient *CiweimaoClient) SetVersion(version string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Version = version
	return ciweimaoClient
}

func (ciweimaoClient *CiweimaoClient) SetDebug() *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.Debug = true
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetDebug(true)
	file, err := os.OpenFile("catapi.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open file error !")
	}
	ciweimaoClient.Ciweimao.Req.FileLog = file
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetProxy(proxy string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetProxy(proxy)
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
func (ciweimaoClient *CiweimaoClient) SetUserAgent(value string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetHeader("User-Agent", value+" "+ciweimaoClient.Ciweimao.Req.Version)
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetRetryCount(retryCount int) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetRetryCount(retryCount)
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetBaseURL(baseURL string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetBaseURL(baseURL)
	return ciweimaoClient
}
func UnescapeUnicode(raw string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

func (ciweimaoClient *CiweimaoClient) SetAccount(account string) *CiweimaoClient {
	if unescapeUnicode, err := UnescapeUnicode(account); err != nil {
		log.Println("set account error", err)
	} else if !strings.Contains(unescapeUnicode, "书客") {
		log.Println("set account error:", "account is not contains 书客")
	} else {
		ciweimaoClient.Ciweimao.Req.Account = unescapeUnicode
	}
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetAuth(account, loginToken string) *CiweimaoClient {
	return ciweimaoClient.SetAccount(account).SetLoginToken(loginToken)
}
