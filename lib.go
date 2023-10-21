package catapi

import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/catnovelapi/catapi/options"
)

func NewCiweimaoClient(options ...options.CiweimaoOption) *catapi.Ciweimao {
	client := &catapi.Ciweimao{
		Host:        "https://app.hbooker.com",
		Version:     "2.9.290",
		DecodeKey:   "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
		DeviceToken: "ciweimao_",
	}
	for _, option := range options {
		option.Apply(client)
	}

	client.Headers = map[string]any{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + client.Version,
	}
	return client
}
