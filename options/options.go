package options

import "github.com/catnovelapi/catapi/catapi"

type CiweimaoOption interface {
	Apply(*catapi.Ciweimao)
}
type CiweimaoOptionFunc func(*catapi.Ciweimao)

func (optionFunc CiweimaoOptionFunc) Apply(c *catapi.Ciweimao) {
	optionFunc(c)
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
func Debug() CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Debug = true
	})
}
func Proxy(proxy string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		c.Proxy = proxy
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
func Auth(account, loginToken string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *catapi.Ciweimao) {
		Account(account).Apply(c)
		LoginToken(loginToken).Apply(c)
	})
}
