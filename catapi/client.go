package catapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func (cat *Ciweimao) BuilderClient() *resty.Request {
	restyClient := resty.New().SetRetryCount(5).SetProxy(cat.Proxy)
	return restyClient.R().SetDebug(cat.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"User-Agent":   "Android com.kuangxiangciweimao.novel " + cat.Version,
		}).
		SetFormData(map[string]string{
			"device_token": "ciweimao_",
			"app_version":  cat.Version,
			"login_token":  cat.LoginToken,
			"account":      cat.Account,
		})
}
func (cat *Ciweimao) PostAPI(url string, data map[string]string) (gjson.Result, error) {
	response, err := cat.Post(url, data)
	if err != nil {
		return gjson.Result{}, err
	}
	decodeText, err := cat.DecodeEncryptText(response.String(), cat.DecodeKey)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.Parse(decodeText), nil
}
func (cat *Ciweimao) GetAPI(url string, data map[string]string) (gjson.Result, error) {
	response, err := cat.Get(url, data)
	if err != nil {
		return gjson.Result{}, err
	}
	decodeText, err := cat.DecodeEncryptText(response.String(), cat.DecodeKey)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.Parse(decodeText), nil
}

func (cat *Ciweimao) Post(url string, data map[string]string) (*resty.Response, error) {
	if data == nil {
		data = map[string]string{}
	}
	return cat.BuilderClient().SetFormData(data).Post("https://app.hbooker.com" + url)
}

func (cat *Ciweimao) Get(url string, data map[string]string) (*resty.Response, error) {
	if data == nil {
		data = map[string]string{}
	}
	return cat.BuilderClient().SetFormData(data).Get("https://app.hbooker.com" + url)
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

func (cat *Ciweimao) DecodeEncryptText(str string, decodeKey string) (string, error) {
	if decodeKey == "" {
		return str, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	raw, err := aesDecrypt(decodeKey, decoded)
	if err != nil {
		return "", err
	}
	if len(raw) == 0 {
		return "", errors.New("解密内容为空,请检查解密内容内容是否正确")
	}
	return string(raw), nil
}
