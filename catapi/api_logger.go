package catapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"strings"
)

type ApiLogger struct {
	c        *CiweimaoRequest
	builder  *strings.Builder
	response *resty.Response
}

func NewApiLogger(response *resty.Response, c *CiweimaoRequest) *ApiLogger {
	return &ApiLogger{c: c, response: response, builder: &strings.Builder{}}
}

func (apiLogger *ApiLogger) addLogger(err error) {
	if !apiLogger.c.Debug {
		return
	}
	apiLogger.builder.WriteString("\nResponse Info:\n")
	if err != nil {
		apiLogger.builder.WriteString(fmt.Sprintf("  Error: %s\n", err.Error()))
	}
	apiLogger.addStatus()
	apiLogger.addHeader()
	apiLogger.addCookies()
	apiLogger.addForm()
	apiLogger.addResult()

	if err = apiLogger.saveLogToFile(); err != nil {
		log.Println(err)
	}
}

func (apiLogger *ApiLogger) saveLogToFile() error {
	apiLogger.builder.WriteString("============================================================END\n")
	_, err := apiLogger.c.FileLog.WriteString(apiLogger.builder.String())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (apiLogger *ApiLogger) addStatus() {
	apiLogger.builder.WriteString(fmt.Sprintf(
		"  Status Code: %d\n  Status : %s\n  Proto: %s\n  Time: %s\n  Received At: %s\n",
		apiLogger.response.StatusCode(),
		apiLogger.response.Status(),
		apiLogger.response.Proto(),
		apiLogger.response.Time(),
		apiLogger.response.Time(),
	))
}
func (apiLogger *ApiLogger) addHeader() {
	if len(apiLogger.response.Header()) > 0 {
		apiLogger.builder.WriteString("  Header:\n")
		for k, v := range apiLogger.response.Header() {
			apiLogger.builder.WriteString(fmt.Sprintf("    Header     : %s=%s\n", k, v))
		}
	}
}

func (apiLogger *ApiLogger) addCookies() {
	if len(apiLogger.response.Cookies()) > 0 {
		apiLogger.builder.WriteString("  Cookies:\n")
		for _, cookie := range apiLogger.response.Cookies() {
			apiLogger.builder.WriteString(fmt.Sprintf("    Cookie     : %s=%s\n", cookie.Name, cookie.Value))
		}
	}
}

func (apiLogger *ApiLogger) addForm() {
	if apiLogger.response.Request.FormData != nil {
		apiLogger.builder.WriteString("  Form:\n")
		for k, v := range apiLogger.response.Request.FormData {
			apiLogger.builder.WriteString(fmt.Sprintf("    Form       : %s=%s\n", k, v))
		}
	}
}
func (apiLogger *ApiLogger) addResult() {
	result := string(apiLogger.response.Body())
	if result == "" {
		apiLogger.builder.WriteString(fmt.Sprintf("  Result       :\n %s\n", "empty"))
		return
	}
	var err error
	var jsonString string
	if !gjson.Valid(result) {
		jsonString, err = apiLogger.c.DecodeEncryptText(result, decodeKey)
		if err != nil {
			apiLogger.builder.WriteString(fmt.Sprintf("  Decode Error: %s\n", err.Error()))
			apiLogger.builder.WriteString(fmt.Sprintf("  Result       :\n %s\n", result))
			return
		}
	}
	apiLogger.builder.WriteString(fmt.Sprintf("  Result       :\n %s\n", IndentJson(jsonString)))
}
func IndentJson(a string) string {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal([]byte(a), &objmap)
	if err != nil {
		log.Println(err)
		return a + "\n" + err.Error()
	}
	formatted, err := json.MarshalIndent(objmap, "", "  ")
	if err != nil {
		return a + "\n" + err.Error()
	}
	return string(formatted)
}
