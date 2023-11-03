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

type CiweimaoClient struct {
	Ciweimao *catapi.Ciweimao
}

func NewCiweimaoClient(options ...options.CiweimaoRequestOption) *CiweimaoClient {
	client := &CiweimaoClient{&catapi.Ciweimao{
		Req: &catapi.CiweimaoRequest{
			BuilderClient: resty.New().SetRetryCount(5),
			Debug:         false, Version: "2.9.290"},
	}}
	for _, option := range options {

		option.Apply(client.Ciweimao.Req)
	}
	if client.Ciweimao.Req.Debug {
		client.Ciweimao.Req.BuilderClient.SetDebug(client.Ciweimao.Req.Debug)
		file, err := os.OpenFile("catapi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !")
		}
		client.Ciweimao.Req.FileLog = file
	}
	if client.Ciweimao.Req.Proxy != "" {
		client.Ciweimao.Req.BuilderClient.SetProxy(client.Ciweimao.Req.Proxy)
	}
	client.Ciweimao.Req.BuilderClient.SetFormData(map[string]string{
		"device_token": deviceToken,
		"app_version":  client.Ciweimao.Req.Version,
		"login_token":  client.Ciweimao.Req.LoginToken,
		"account":      client.Ciweimao.Req.Account,
	})

	client.Ciweimao.Req.BuilderClient.SetHeaders(map[string]string{"User-Agent": useragent + client.Ciweimao.Req.Version})
	return client
}
