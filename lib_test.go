package catapi

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var client *CiweimaoClient

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client = NewCiweimaoClient().
		SetDebug().
		SetProxy(os.Getenv("PROXY")).
		SetAuth(os.Getenv("CAT_ACCOUNT"), os.Getenv("CAT_LOGIN_TOKEN"))
}
func TestNewCiweimaoSearchBooks(t *testing.T) {
	searchByKeywordApi, err := client.Ciweimao.SearchByKeywordApi(os.Getenv("SEARCH_KEYWORD"), "0")
	if err != nil {
		t.Error(err)
		return
	}
	for _, book := range searchByKeywordApi.Get("data").Array() {
		println(book.Get("book_id").String())
		println(book.Get("book_name").String())
	}
}

func TestCiweimaoBookInfo(t *testing.T) {
	bookInfo, err := client.Ciweimao.BookInfoApiByBookId(os.Getenv("BOOK_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(bookInfo.String())
}
