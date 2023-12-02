package catapi

import (
	"github.com/catnovelapi/builder"
	"github.com/catnovelapi/catapi/catapi"
	"log"
	"strconv"
	"strings"
)

type CiweimaoClient struct {
	defBaseURL       string
	defaultUserAgent string
	defaultVersion   string
	Ciweimao         *catapi.Ciweimao
}

func NewCiweimaoClient() *CiweimaoClient {
	client := &CiweimaoClient{
		defaultVersion:   "2.9.290",
		defBaseURL:       "https://app.hbooker.com",
		defaultUserAgent: "Android com.kuangxiangciweimao.novel",
		Ciweimao: &catapi.Ciweimao{
			Req: &catapi.CiweimaoRequest{BuilderClient: builder.NewClient()},
		},
	}
	client.SetRetryCount(7).
		SetBaseURL(client.defBaseURL).
		SetVersion(client.defaultVersion).
		SetUserAgent(client.defaultUserAgent).
		SetDeviceToken("ciweimao_")

	if version, err := client.Ciweimao.GetVersionApi(); err == nil {
		client.SetVersion(version)
	}
	// refresh user agent
	client.SetUserAgent(client.defaultUserAgent)
	return client
}
func (ciweimaoClient *CiweimaoClient) SetDeviceToken(deviceToken string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetHeader("device_token", deviceToken)
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetVersion(version string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetHeader("app_version", version)
	return ciweimaoClient
}

func (ciweimaoClient *CiweimaoClient) SetDebug() *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetDebug()
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetDebugFile("catapi.txt")
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
		ciweimaoClient.Ciweimao.Req.BuilderClient.SetHeader("login_token", loginToken)
	}
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetUserAgent(value string) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetUserAgent(value + " " + ciweimaoClient.Ciweimao.Req.Version)
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetRetryCount(retryCount int) *CiweimaoClient {
	ciweimaoClient.Ciweimao.Req.BuilderClient.SetRetryNumber(retryCount)
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
		ciweimaoClient.Ciweimao.Req.BuilderClient.SetHeader("account", unescapeUnicode)
	}
	return ciweimaoClient
}
func (ciweimaoClient *CiweimaoClient) SetAuth(account, loginToken string) *CiweimaoClient {
	return ciweimaoClient.SetAccount(account).SetLoginToken(loginToken)
}
