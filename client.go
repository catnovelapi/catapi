package catapi

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/catnovelapi/builder"
	"github.com/catnovelapi/catapi/catapi/decrypt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
)

type ciweimaoAuthentication struct {
	Account     string `json:"account"`
	LoginToken  string `json:"login_token"`
	DeviceToken string `json:"device_token"`
	Version     string `json:"app_version"`
}

type Client struct {
	m              sync.RWMutex // 用于保证线程安全
	debug          bool
	retryCount     int
	baseURL        string
	userAgent      string
	proxy          string
	authentication ciweimaoAuthentication
}

type API struct {
	client        *Client
	builderClient *builder.Client
}

func NewClient() *Client {
	return &Client{
		retryCount: 7,
		baseURL:    "https://app.hbooker.com",
		userAgent:  "Android com.kuangxiangciweimao.novel ",
		proxy:      "",
		authentication: ciweimaoAuthentication{
			Version:     "2.9.290",
			DeviceToken: "ciweimao_",
		},
	}
}

// StructToMap converts a CiweimaoAuthentication struct to a map[string]interface{}
func structToMap(auth any) (map[string]interface{}, error) {
	// 序列化结构体为JSON
	jsonBytes, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}

	// 反序列化JSON到map
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// R 方法用于实例化一些默认的参数, 并返回一个 Client 对象的指针。
func (client *Client) R() *API {
	builderClient := builder.NewClient().
		SetBaseURL(client.baseURL).
		SetRetryCount(client.retryCount).
		SetUserAgent(client.userAgent + client.authentication.Version).
		SetResultFunc(decrypt.DecodeFunc)

	if client.proxy != "" {
		builderClient.SetProxy(client.proxy)
	}
	authMap, err := structToMap(client.authentication)
	if err != nil {
		fmt.Println(err)
	} else {
		builderClient.SetQueryParams(authMap)
	}

	return &API{client: client, builderClient: builderClient}
}

//// UpdateVersion 方法用于更新版本号, 它会调用 API 的 GetVersionApi 方法, 并将返回的版本号设置为 HTTP 请求的版本号。
//func (client *Client) UpdateVersion() *Client {
//	if version, err := client.Ciweimao.GetVersionApi(); err == nil {
//		client.SetVersion(version)
//		client.SetUserAgent("Android com.kuangxiangciweimao.novel " + version)
//	}
//	// refresh user agent
//	return client
//}

// SetDeviceToken 方法用于设置 HTTP 请求的设备号。它接收一个 string 类型的参数，该参数表示设备号的值。
func (client *Client) SetDeviceToken(deviceToken string) *Client {
	client.authentication.DeviceToken = deviceToken
	return client
}

// SetVersion 方法用于设置 HTTP 请求的版本号。它接收一个 string 类型的参数，该参数表示版本号的值。
func (client *Client) SetVersion(version string) *Client {
	client.authentication.Version = version
	return client
}

// SetDebug 方法用于设置是否输出调试信息。它接收一个 bool 类型的参数，该参数表示是否输出调试信息。
func (client *Client) SetDebug() *Client {
	client.debug = true
	return client
}

// SetProxy	方法用于设置 HTTP 请求的代理。它接收一个 string 类型的参数，该参数表示代理的值。
func (client *Client) SetProxy(proxy string) *Client {
	client.proxy = proxy
	return client
}

// SetLoginToken 方法用于设置 HTTP 请求的登录令牌。它接收一个 string 类型的参数，该参数表示登录令牌的值。
func (client *Client) SetLoginToken(loginToken string) *Client {
	if len(loginToken) != 32 {
		log.Println("loginToken length is not 32")
	} else {
		client.authentication.LoginToken = loginToken
	}
	return client
}

// SetUserAgent 方法用于设置 HTTP 请求的 User-Agent 部分。它接收一个 string 类型的参数，该参数表示 User-Agent 的值。
func (client *Client) SetUserAgent(value string) *Client {
	client.userAgent = value
	return client
}

// SetRetryCount 方法用于设置重试次数。它接收一个 int 类型的参数，该参数表示重试次数。
func (client *Client) SetRetryCount(retryCount int) *Client {
	client.retryCount = retryCount
	return client
}

// SetBaseURL 方法用于设置 HTTP 请求的 BaseURL 部分。它接收一个 string 类型的参数，该参数表示 BaseURL 的值。
func (client *Client) SetBaseURL(baseURL string) *Client {
	client.baseURL = baseURL
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
		//client.API.Req.BuilderClient.SetQueryParam("account", unescapeUnicode)
		client.authentication.Account = unescapeUnicode
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
