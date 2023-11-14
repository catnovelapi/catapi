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
	for _, book := range searchByKeywordApi.Array() {
		println(book.Get("book_id").String())
		println(book.Get("book_name").String())
	}
}
func TestUserInfo(t *testing.T) {
	accountInfo, err := client.Ciweimao.AccountInfoApi()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(accountInfo.String())
}

func TestChapterList(t *testing.T) {
	chaptersCatalog, err := client.Ciweimao.ChaptersCatalogV2Api(os.Getenv("BOOK_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(chaptersCatalog.String())
}
func TestCiweimaoBookInfo(t *testing.T) {
	bookInfo, err := client.Ciweimao.BookInfoApiByBookId(os.Getenv("BOOK_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(bookInfo.String())
}

func TestCiweimaoChapterInfo(t *testing.T) {
	chapterInfo, err := client.Ciweimao.ChapterInfoApi(os.Getenv("CHAPTER_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(chapterInfo.String())
}

func TestCiweimaoBookShelfIdList(t *testing.T) {
	api, err := client.Ciweimao.BookShelfIdListApi()
	if err != nil {
		return
	}
	fmt.Println(api.String())
}
func TestCiweimaoBookShelfList(t *testing.T) {
	api, err := client.Ciweimao.BookShelfListApi(os.Getenv("SHELF_ID"))
	if err != nil {
		return
	}
	fmt.Println(api.String())
}
func TestCiweimaoReviewList(t *testing.T) {
	api, err := client.Ciweimao.ReviewListApi(os.Getenv("BOOK_ID"), "0")
	if err != nil {
		return
	}
	fmt.Println(api.String())
}

func TestCiweimaoReviewCommentListApi(t *testing.T) {
	api, err := client.Ciweimao.ReviewCommentListApi(os.Getenv("REVIEW_ID"), "0")
	if err != nil {
		return
	}
	fmt.Println(api.String())
}

func TestCiweimaoReviewCommentReplyListApi(t *testing.T) {
	api, err := client.Ciweimao.ReviewCommentReplyListApi(os.Getenv("COMMENT_ID"), "0")
	if err != nil {
		return
	}
	fmt.Println(api.String())
}
