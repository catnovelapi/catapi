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
	client := NewCiweimaoClient(
		options.Debug(),
		options.Proxy(os.Getenv("PROXY")),
		options.Auth(os.Getenv("CAT_ACCOUNT"), os.Getenv("CAT_LOGIN_TOKEN")),
	)
	searchByKeywordApi, err := client.SearchByKeywordApi(os.Getenv("SEARCH_KEYWORD"), "0")
	if err != nil {
		t.Error(err)
		return
	}
	for _, book := range searchByKeywordApi.Get("data").Array() {
		println(book.Get("book_id").String())
		println(book.Get("book_name").String())
	}
}
