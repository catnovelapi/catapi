package options

import "github.com/catnovelapi/catapi/catapi"

type CiweimaoRequestOption interface {
	Apply(*catapi.CiweimaoRequest)
}
type CiweimaoRequestOptionFunc func(*catapi.CiweimaoRequest)

func (optionFunc CiweimaoRequestOptionFunc) Apply(c *catapi.CiweimaoRequest) {
	optionFunc(c)
}

func Version(version string) CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		c.Version = version
	})
}
func Debug() CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		c.Debug = true
	})
}
func Proxy(proxy string) CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		c.Proxy = proxy
	})
}
func LoginToken(loginToken string) CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		c.LoginToken = loginToken
	})
}
func Account(account string) CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		c.Account = account
	})
}
func Auth(account, loginToken string) CiweimaoRequestOption {
	return CiweimaoRequestOptionFunc(func(c *catapi.CiweimaoRequest) {
		Account(account).Apply(c)
		LoginToken(loginToken).Apply(c)
	})
}
