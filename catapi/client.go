package catapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"os"
)

type CiweimaoRequest struct {
	Debug         bool
	FileLog       *os.File
	Proxy         string
	Host          string
	Version       string
	LoginToken    string
	Account       string
	BuilderClient *resty.Client
}

const decodeKey = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"

func (request *CiweimaoRequest) addLogger(resp *resty.Response, err error) {
	if !request.Debug {
		return
	}
	var responseInfo string
	responseInfo = "Response Info:\n"
	if err != nil {
		responseInfo += fmt.Sprintf("  Error: %s\n", err.Error())
	}
	responseInfo += fmt.Sprintf("  Status Code: %d\n", resp.StatusCode())
	responseInfo += fmt.Sprintf("  Status : %s\n", resp.Status())
	responseInfo += fmt.Sprintf("  Proto      :%s\n", resp.Proto())
	responseInfo += fmt.Sprintf("  Time	   :%s\n", resp.Time())
	responseInfo += fmt.Sprintf("  Received At:%s\n", resp.Time())
	if len(resp.Header()) > 0 {
		responseInfo += fmt.Sprintf("  Header:\n")
		for k, v := range resp.Header() {
			responseInfo += fmt.Sprintf("    Header     : %s=%s\n", k, v)
		}
	}
	if len(resp.Cookies()) > 0 {
		responseInfo += fmt.Sprintf("  Cookies:\n")
		for _, cookie := range resp.Cookies() {
			responseInfo += fmt.Sprintf("    Cookie     : %s=%s\n", cookie.Name, cookie.Value)
		}
	}
	if resp.Request.FormData != nil {
		responseInfo += fmt.Sprintf("  Form:\n")
		for k, v := range resp.Request.FormData {
			responseInfo += fmt.Sprintf("    Form       : %s=%s\n", k, v)
		}
	}
	result := string(resp.Body())
	if result != "" {
		if gjson.Valid(result) {
			responseInfo += fmt.Sprintf("  Result       :\n %s\n", result)
		} else {
			result, err = request.DecodeEncryptText(result, decodeKey)
			if err != nil {
				responseInfo += fmt.Sprintf("  Decode Error: %s\n", err.Error())
				responseInfo += fmt.Sprintf("  Result       :\n %s\n", result)
			} else {
				responseInfo += fmt.Sprintf("  Result       :\n %s\n", result)
			}
		}
	}
	responseInfo += fmt.Sprintf("============================================================\n")
	_, err = request.FileLog.WriteString(responseInfo)
	if err != nil {
		log.Println(err)
		return
	}
}
func (request *CiweimaoRequest) PostAPI(url string, data map[string]string) (gjson.Result, error) {
	if data == nil {
		data = map[string]string{}
	}
	response, err := request.BuilderClient.R().SetFormData(data).Post(baseUrl + url)
	defer request.addLogger(response, err)
	if err != nil {
		return gjson.Result{}, err
	}
	if response.StatusCode() != 200 {
		return gjson.Result{}, errors.New("status error: " + response.Status())
	}
	responseText := response.String()
	if responseText == "" {
		return gjson.Result{}, errors.New("responseText is empty, please check your network")
	}
	if !gjson.Valid(responseText) {
		responseText, err = request.DecodeEncryptText(response.String(), decodeKey)
		if err != nil {
			return gjson.Result{}, err
		}
	}
	return gjson.Parse(responseText), nil
}
func (request *CiweimaoRequest) GetAPI(url string, data map[string]string) (gjson.Result, error) {
	if data == nil {
		data = map[string]string{}
	}
	response, err := request.BuilderClient.R().SetFormData(data).Get(baseUrl + url)
	defer request.addLogger(response, err)
	if err != nil {
		return gjson.Result{}, err
	}
	if response.StatusCode() != 200 {
		return gjson.Result{}, errors.New("status error: " + response.Status())
	}
	responseText := response.String()
	if responseText == "" {
		return gjson.Result{}, errors.New("responseText is empty, please check your network")
	}
	if !gjson.Valid(responseText) {
		responseText, err = request.DecodeEncryptText(response.String(), decodeKey)
		if err != nil {
			return gjson.Result{}, err
		}
	}
	return gjson.Parse(responseText), nil
}

var IV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SHA256 sha256 编码
func SHA256(data []byte) []byte {
	ret := sha256.Sum256(data)
	return ret[:]
}

// LoadKey 读取解密密钥
func LoadKey(EncryptKey string) []byte {
	Key := SHA256([]byte(EncryptKey))
	return Key[:32]
}

func aesDecrypt(EncryptKey string, ciphertext []byte) ([]byte, error) {
	key := LoadKey(EncryptKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, IV)
	plainText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plainText, ciphertext)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}

// PKCS7UnPadding 对齐
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])
	return plainText[:(length - unpadding)]
}

func (request *CiweimaoRequest) DecodeEncryptText(str string, decodeKey string) (string, error) {
	if decodeKey == "" {
		return "", errors.New("解密密钥为空,请检查解密密钥是否正确")
	}
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	raw, err := aesDecrypt(decodeKey, decoded)
	if err != nil {
		return "", errors.New("解密失败,请检查解密密钥是否正确")
	}
	if len(raw) == 0 {
		return "", errors.New("解密内容为空,请检查解密内容内容是否正确")
	}
	return string(raw), nil
}
