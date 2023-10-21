package catapi

import (
	"github.com/catnovelapi/catapi/options"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func init() {
}
func TestNewCiweimaoClient(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CiweimaoClient := NewCiweimaoClient(
		options.Debug(),
		options.Proxy(os.Getenv("PROXY")),
		options.Auth(os.Getenv("CAT_ACCOUNT"), os.Getenv("CAT_LOGIN_TOKEN")),
	)
	searchByKeywordApi, err := CiweimaoClient.SearchByKeywordApi("", "1", "0")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(searchByKeywordApi)
}
