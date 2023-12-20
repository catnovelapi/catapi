package catapi

import (
	"fmt"
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

	if result := gjson.Parse(response.String()); result.Get("code").String() != "100000" {
		return result, fmt.Errorf("response error: %s", result.Get("tip").String())
	} else {
		return result, nil
	}
}
