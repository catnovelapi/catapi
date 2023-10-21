package options

import "github.com/catnovelapi/catapi/catapi"

type HttpClient struct {
	Debug     bool
	MaxRetry  int
	DecodeKey string
}

type HttpOption interface {
	Apply(*HttpClient)
}
type HttpOptionFunc func(*HttpClient)

func (optionFunc HttpOptionFunc) Apply(c *HttpClient) {
	optionFunc(c)
}
func Debug() HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.Debug = true
	})
}
func NoDecode() HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.DecodeKey = ""
	})
}
func MaxRetry(retry int) HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.MaxRetry = retry
	})
}

type CiweimaoOption interface {
	Apply(*catapi.Ciweimao)
}
type CiweimaoOptionFunc func(*catapi.Ciweimao)

func (optionFunc CiweimaoOptionFunc) Apply(c *catapi.Ciweimao) {
	optionFunc(c)
}

func ApiBase(host string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Host = host
	})
}
func Version(version string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Version = version
	})
}
func DecodeKey(decodeKey string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.DecodeKey = decodeKey
	})
}
func DeviceToken(deviceToken string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.DeviceToken = deviceToken
	})
}
func AppVersion(appVersion string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Version = appVersion
	})
}
func LoginToken(loginToken string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.LoginToken = loginToken
	})
}
func Account(account string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Account = account
	})
}
