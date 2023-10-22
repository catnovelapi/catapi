package catapi

import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/catnovelapi/catapi/options"
)

func NewCiweimaoClient(options ...options.CiweimaoOption) *catapi.Ciweimao {
	client := &catapi.Ciweimao{
		Debug:   false,
		Version: "2.9.290",
	}
	for _, option := range options {
		option.Apply(client)
	}
	return client.Builder()
}
