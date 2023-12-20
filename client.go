package catapi

import (
	"crypto/rand"
	"fmt"
	"github.com/catnovelapi/builder"
	"github.com/catnovelapi/catapi/catapi/decrypt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	m        sync.RWMutex // 用于保证线程安全
	Ciweimao *Ciweimao    // 指向 Ciweimao 的指针, 用于调用 Ciweimao 的方法和接口
}

func NewCiweimaoClient() *Client {
	builderClient := builder.NewClient().SetResultFunc(func(result string) (string, error) {
		text, err := decrypt.DecodeEncryptText(result, "")
		if err != nil {
			return "", err
		}
		return text, nil
	})
	builderClient.SetContentType("application/x-www-form-urlencoded")
	client := &Client{Ciweimao: &Ciweimao{
		Req: &CiweimaoRequest{BuilderClient: builderClient},
	}}

	return client
}

// R 方法用于实例化一些默认的参数, 并返回一个 Client 对象的指针。
func (client *Client) R() *Client {
	return client.
		SetRetryCount(7).
		SetVersion("2.9.290").
		SetDeviceToken("ciweimao_").
		SetBaseURL("https://app.hbooker.com").
		SetUserAgent("Android com.kuangxiangciweimao.novel 2.9.290")
}

// UpdateVersion 方法用于更新版本号, 它会调用 Ciweimao 的 GetVersionApi 方法, 并将返回的版本号设置为 HTTP 请求的版本号。
func (client *Client) UpdateVersion() *Client {
	if version, err := client.Ciweimao.GetVersionApi(); err == nil {
		client.SetVersion(version)
		client.SetUserAgent("Android com.kuangxiangciweimao.novel " + version)
	}
	// refresh user agent
	return client
}

// SetDeviceToken 方法用于设置 HTTP 请求的设备号。它接收一个 string 类型的参数，该参数表示设备号的值。
func (client *Client) SetDeviceToken(deviceToken string) *Client {
	client.Ciweimao.Req.BuilderClient.SetQueryParam("device_token", deviceToken)
	return client
}

// SetVersion 方法用于设置 HTTP 请求的版本号。它接收一个 string 类型的参数，该参数表示版本号的值。
func (client *Client) SetVersion(version string) *Client {
	client.Ciweimao.Req.BuilderClient.SetQueryParam("app_version", version)
	return client
}

// SetDebug 方法用于设置是否输出调试信息。它接收一个 bool 类型的参数，该参数表示是否输出调试信息。
func (client *Client) SetDebug() *Client {
	client.Ciweimao.Req.BuilderClient.SetDebug()
	client.Ciweimao.Req.BuilderClient.SetDebugFile("catapi")
	return client
}

// SetProxy	方法用于设置 HTTP 请求的代理。它接收一个 string 类型的参数，该参数表示代理的值。
func (client *Client) SetProxy(proxy string) *Client {
	client.Ciweimao.Req.BuilderClient.SetProxy(proxy)
	return client
}

// SetLoginToken 方法用于设置 HTTP 请求的登录令牌。它接收一个 string 类型的参数，该参数表示登录令牌的值。
func (client *Client) SetLoginToken(loginToken string) *Client {
	if len(loginToken) != 32 {
		log.Println("loginToken length is not 32")
	} else {
		client.Ciweimao.Req.BuilderClient.SetQueryParam("login_token", loginToken)
	}
	return client
}

// SetUserAgent 方法用于设置 HTTP 请求的 User-Agent 部分。它接收一个 string 类型的参数，该参数表示 User-Agent 的值。
func (client *Client) SetUserAgent(value string) *Client {
	client.Ciweimao.Req.BuilderClient.SetUserAgent(value)
	return client
}

// SetRetryCount 方法用于设置重试次数。它接收一个 int 类型的参数，该参数表示重试次数。
func (client *Client) SetRetryCount(retryCount int) *Client {
	client.Ciweimao.Req.BuilderClient.SetRetryCount(retryCount)
	return client
}

// SetBaseURL 方法用于设置 HTTP 请求的 BaseURL 部分。它接收一个 string 类型的参数，该参数表示 BaseURL 的值。
func (client *Client) SetBaseURL(baseURL string) *Client {
	client.Ciweimao.Req.BuilderClient.SetBaseURL(baseURL)
	return client
}

// UnescapeUnicode 方法用于将 Unicode 编码的字符串转换为中文字符串。它接收一个 string 类型的参数，该参数表示 Unicode 编码的字符串。
func UnescapeUnicode(raw string) (string, error) {
	// strconv.Unquote 方法用于将字符串中的转义字符转换为相应的字符
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

// SetAccount 方法用于设置 HTTP 请求的账号。它接收一个 string 类型的参数，该参数表示账号的值。
func (client *Client) SetAccount(account string) *Client {
	if unescapeUnicode, err := UnescapeUnicode(account); err != nil {
		log.Println("set account error", err)
	} else if !strings.Contains(unescapeUnicode, "书客") {
		log.Println("set account error:", "account is not contains 书客")
	} else {
		client.Ciweimao.Req.BuilderClient.SetQueryParam("account", unescapeUnicode)
	}
	return client
}

// SetAuthentication 方法用于设置 HTTP 请求的账号和登录令牌。它接收两个 string 类型的参数，第一个参数表示账号的值，第二个参数表示登录令牌的值。
func (client *Client) SetAuthentication(account, loginToken string) *Client {
	return client.SetAccount(account).SetLoginToken(loginToken)
}

func (client *Client) AndroidID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		log.Fatal(err)
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
