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
	"os"
)

type CiweimaoRequest struct {
	Debug         bool
	FileLog       *os.File
	Version       string
	LoginToken    string
	Account       string
	BuilderClient *resty.Client
}

const decodeKey = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"

func (request *CiweimaoRequest) getDefaultAuthenticationFormData() map[string]string {
	return map[string]string{
		"device_token": "ciweimao_",
		"app_version":  request.Version,
		"login_token":  request.LoginToken,
		"account":      request.Account,
	}
}
func (request *CiweimaoRequest) PostAPI(url string, data map[string]string) (gjson.Result, error) {
	formData := request.getDefaultAuthenticationFormData()
	if data != nil {
		for k, v := range data {
			formData[k] = v
		}
	}
	response, err := request.BuilderClient.R().SetFormData(formData).Post(url)
	defer NewApiLogger(response, request).addLogger(err)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("request error: %s", err.Error())
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
			return gjson.Result{}, fmt.Errorf("decode error: %s", err.Error())
		}
	}
	return gjson.Parse(responseText), nil
}
func (request *CiweimaoRequest) GetAPI(url string, data map[string]string) (gjson.Result, error) {
	formData := request.getDefaultAuthenticationFormData()
	if data != nil {
		for k, v := range data {
			formData[k] = v
		}
	}
	response, err := request.BuilderClient.R().SetFormData(formData).Get(url)
	defer NewApiLogger(response, request).addLogger(err)
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
		return "", fmt.Errorf("base64 decode error: %s", err.Error())
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
