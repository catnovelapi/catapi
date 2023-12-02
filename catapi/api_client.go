package catapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/catnovelapi/builder"
	"github.com/tidwall/gjson"
)

type CiweimaoRequest struct {
	Version       string
	Account       string
	BuilderClient *builder.Client
}

func (request *CiweimaoRequest) Post(url string, data map[string]any) (gjson.Result, error) {
	req := request.BuilderClient.R()
	if data != nil {
		req.SetQueryParams(data)
	}
	response, err := req.Post(url)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("request error: %s", err.Error())
	}
	responseText := response.String()
	if responseText == "" {
		return gjson.Result{}, errors.New("responseText is empty, please check your network")
	}
	if !gjson.Valid(responseText) {
		responseText, err = request.DecodeEncryptText(response.String(), "")
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

// SHA256 sha256 编码
func SHA256(data []byte) []byte {
	ret := sha256.Sum256(data)
	return ret[:]
}

func aesDecrypt(EncryptKey string, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(SHA256([]byte(EncryptKey))[:32])
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
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
		decodeKey = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
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
