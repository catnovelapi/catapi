package catapi

import (
	"fmt"
	"github.com/catnovelapi/catapi/catapi/decrypt"
	"github.com/tidwall/gjson"
)

func (cat *API) post(url string, data map[string]string) (gjson.Result, error) {
	req := cat.builderClient.R()
	if data != nil {
		req.SetQueryParams(data)
	}
	response, err := req.Post(url)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("request error: %s", err.Error())
	}
	var responseText = response.String()
	if !gjson.Valid(responseText) {
		responseText, err = decrypt.DecodeEncryptText(responseText, "")
		if err != nil {
			return gjson.Result{}, fmt.Errorf("decode error: %s", err.Error())
		}
	}
	gjsonResponseText := gjson.Parse(responseText)
	if gjsonResponseText.Get("code").String() != "100000" {
		return gjson.Result{}, fmt.Errorf("response error: %s", gjsonResponseText.Get("tip").String())
	}
	return gjsonResponseText, nil
}
