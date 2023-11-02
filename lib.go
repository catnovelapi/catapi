package catapi

import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/catnovelapi/catapi/options"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

const deviceToken = "ciweimao_"
const useragent = "Android com.kuangxiangciweimao.novel "

func NewCiweimaoClient(options ...options.CiweimaoOption) *catapi.Ciweimao {
	client := &catapi.Ciweimao{Debug: false, Version: "2.9.290"}
	for _, option := range options {
		option.Apply(client)
	}
	client.BuilderClient = resty.New().SetRetryCount(5)
	if client.Debug {
		client.BuilderClient.SetDebug(client.Debug)
		file, err := os.OpenFile("catapi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !")
		}
		client.FileLog = file
	}
	if client.Proxy != "" {
		client.BuilderClient.SetProxy(client.Proxy)
	}
	client.BuilderClient.SetFormData(map[string]string{
		"device_token": deviceToken,
		"app_version":  client.Version,
		"login_token":  client.LoginToken,
		"account":      client.Account,
	})
	client.BuilderClient.SetHeaders(map[string]string{"User-Agent": useragent + client.Version})
	return client
}
