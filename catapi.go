package catapi

import (
	"github.com/catnovelapi/builder"
	"github.com/catnovelapi/catapi/catapi"
	"log"
	"strconv"
	"strings"
)

type Client struct {
	Ciweimao *catapi.Ciweimao
}

func NewCiweimaoClient() *Client {
	builderClient := builder.NewClient()
	client := &Client{Ciweimao: &catapi.Ciweimao{
		Req: &catapi.CiweimaoRequest{BuilderClient: builderClient},
	}}
	return client
}
func (client *Client) R() *Client {
	return client.
		SetRetryCount(7).
		SetVersion("2.9.290").
		SetDeviceToken("ciweimao_").
		SetBaseURL("https://app.hbooker.com").
		SetUserAgent("Android com.kuangxiangciweimao.novel")
}
func (client *Client) UpdateVersion() *Client {
	if version, err := client.Ciweimao.GetVersionApi(); err == nil {
		client.SetVersion(version)
	}
	// refresh user agent
	client.SetUserAgent("Android com.kuangxiangciweimao.novel")
	return client
}
func (client *Client) SetDeviceToken(deviceToken string) *Client {
	client.Ciweimao.Req.BuilderClient.SetHeader("device_token", deviceToken)
	return client
}
func (client *Client) SetVersion(version string) *Client {
	client.Ciweimao.Req.BuilderClient.SetHeader("app_version", version)
	return client
}

func (client *Client) SetDebug() *Client {
	client.Ciweimao.Req.BuilderClient.SetDebug()
	client.Ciweimao.Req.BuilderClient.SetDebugFile("catapi.txt")
	return client
}
func (client *Client) SetProxy(proxy string) *Client {
	client.Ciweimao.Req.BuilderClient.SetProxy(proxy)
	return client
}
func (client *Client) SetLoginToken(loginToken string) *Client {
	if len(loginToken) != 32 {
		log.Println("loginToken length is not 32")
	} else {
		client.Ciweimao.Req.BuilderClient.SetHeader("login_token", loginToken)
	}
	return client
}
func (client *Client) SetUserAgent(value string) *Client {
	client.Ciweimao.Req.BuilderClient.SetUserAgent(value + " " + client.Ciweimao.Req.Version)
	return client
}
func (client *Client) SetRetryCount(retryCount int) *Client {
	client.Ciweimao.Req.BuilderClient.SetRetryNumber(retryCount)
	return client
}
func (client *Client) SetBaseURL(baseURL string) *Client {
	client.Ciweimao.Req.BuilderClient.SetBaseURL(baseURL)
	return client
}
func UnescapeUnicode(raw string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

func (client *Client) SetAccount(account string) *Client {
	if unescapeUnicode, err := UnescapeUnicode(account); err != nil {
		log.Println("set account error", err)
	} else if !strings.Contains(unescapeUnicode, "书客") {
		log.Println("set account error:", "account is not contains 书客")
	} else {
		client.Ciweimao.Req.BuilderClient.SetHeader("account", unescapeUnicode)
	}
	return client
}
func (client *Client) SetAuthentication(account, loginToken string) *Client {
	return client.SetAccount(account).SetLoginToken(loginToken)
}
