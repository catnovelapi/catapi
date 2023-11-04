package catapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"strings"
)

type ApiLogger struct {
	c        *CiweimaoRequest
	response *resty.Response
}

func NewApiLogger(response *resty.Response, c *CiweimaoRequest) *ApiLogger {
	return &ApiLogger{c: c, response: response}
}

func (apiLogger *ApiLogger) addLogger(err error) {
	if !apiLogger.c.Debug {
		return
	}
	var builder strings.Builder
	builder.WriteString("Response Info:\n")
	if err != nil {
		fmt.Fprintf(&builder, "  Error: %s\n", err.Error())
	}
	apiLogger.addStatus(&builder)
	apiLogger.addHeader(&builder)
	apiLogger.addCookies(&builder)
	apiLogger.addForm(&builder)
	apiLogger.addResult(&builder)
	builder.WriteString("============================================================\n")
	_, err = apiLogger.c.FileLog.WriteString(builder.String())
	if err != nil {
		log.Println(err)
		return
	}
}
func (apiLogger *ApiLogger) addStatus(builder *strings.Builder) {
	fmt.Fprintf(builder, "  Status Code: %d\n", apiLogger.response.StatusCode())
	fmt.Fprintf(builder, "  Status : %s\n", apiLogger.response.Status())
	fmt.Fprintf(builder, "  Proto      :%s\n", apiLogger.response.Proto())
	fmt.Fprintf(builder, "  Time	   :%s\n", apiLogger.response.Time())
	fmt.Fprintf(builder, "  Received At:%s\n", apiLogger.response.Time())
}
func (apiLogger *ApiLogger) addHeader(builder *strings.Builder) {
	if len(apiLogger.response.Header()) > 0 {
		builder.WriteString("  Header:\n")
		for k, v := range apiLogger.response.Header() {
			fmt.Fprintf(builder, "    Header     : %s=%s\n", k, v)
		}
	}
}

func (apiLogger *ApiLogger) addCookies(builder *strings.Builder) {
	if len(apiLogger.response.Cookies()) > 0 {
		builder.WriteString("  Cookies:\n")
		for _, cookie := range apiLogger.response.Cookies() {
			fmt.Fprintf(builder, "    Cookie     : %s=%s\n", cookie.Name, cookie.Value)
		}
	}
}

func (apiLogger *ApiLogger) addForm(builder *strings.Builder) {
	if apiLogger.response.Request.FormData != nil {
		builder.WriteString("  Form:\n")
		for k, v := range apiLogger.response.Request.FormData {
			fmt.Fprintf(builder, "    Form       : %s=%s\n", k, v)
		}
	}
}

func (apiLogger *ApiLogger) addResult(builder *strings.Builder) {
	result := string(apiLogger.response.Body())
	if result == "" {
		fmt.Fprintf(builder, "  Result       :\n %s\n", "empty")
		return
	}
	if gjson.Valid(result) {
		fmt.Fprintf(builder, "  Result       :\n %s\n", result)
		return
	}
	result, err := apiLogger.c.DecodeEncryptText(result, decodeKey)
	if err != nil {
		fmt.Fprintf(builder, "  Decode Error: %s\n", err.Error())
		fmt.Fprintf(builder, "  Result       :\n %s\n", result)
	} else {
		fmt.Fprintf(builder, "  Result       :\n %s\n", result)
	}
}
