package catapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"os"
	"regexp"
)

type Ciweimao struct {
	Debug         bool
	FileLog       *os.File
	Proxy         string
	Host          string
	Version       string
	LoginToken    string
	Account       string
	BuilderClient *resty.Client
}

func (cat *Ciweimao) DownloadCover(url string) ([]byte, error) {
	for i := 0; i < 5; i++ {
		response, err := cat.BuilderClient.R().Get(url)
		if err != nil {
			fmt.Printf("download cover error,retry %d times\n", i)
			continue
		}
		return response.Body(), nil
	}
	return nil, fmt.Errorf("download cover error:%s\nurl:%s\n", "retry 5 times", url)
}
func (cat *Ciweimao) AccountInfoApi() (gjson.Result, error) {
	return cat.PostAPI(accountInfoApiPoint, nil)
}

func (cat *Ciweimao) ChaptersCatalogApi(bookId string) (gjson.Result, error) {
	return cat.PostAPI(catalogApiPoint, map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) ChaptersCatalogV2Api(bookId string) (gjson.Result, error) {
	return cat.PostAPI(catalogNewApiPoint, map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) BookInfoApiByBookId(bookId string) (gjson.Result, error) {
	if len(bookId) != 9 {
		return gjson.Result{}, fmt.Errorf("bookId length is not 9")
	}
	return cat.PostAPI(bookInfoApiPoint, map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) BookInfoApiByBookURL(url string) (gjson.Result, error) {
	bookIdMustCompile := regexp.MustCompile(`book/(\d{9})`).FindStringSubmatch(url)
	if len(bookIdMustCompile) < 2 {
		return gjson.Result{}, fmt.Errorf("bookId is empty")
	}
	return cat.BookInfoApiByBookId(bookIdMustCompile[1])
}

func (cat *Ciweimao) SearchByKeywordApi(keyword, page string) (gjson.Result, error) {
	return cat.PostAPI(searchBookApiPoint, map[string]string{"count": "10", "page": page, "category_index": "0", "key": keyword})
}

func (cat *Ciweimao) SearchByTagApi(tagName, page string) (gjson.Result, error) {
	return cat.PostAPI(searchBookApiPoint, map[string]string{"count": "10", "page": page, "category_index": "0", "key": tagName})
}
func (cat *Ciweimao) SignupApi(account string, password string) (gjson.Result, error) {
	return cat.PostAPI(loginApiPoint, map[string]string{"login_name": account, "passwd": password})
}
func (cat *Ciweimao) ChapterCommandApi(chapterId string) (gjson.Result, error) {
	return cat.PostAPI(chapterCommandApiPoint, map[string]string{"chapter_id": chapterId})
}

func (cat *Ciweimao) ChapterInfoApi(chapterId string, command string) (gjson.Result, error) {
	if command == "" || len(command) > 100 {
		return gjson.Result{}, fmt.Errorf("command is empty or too long")
	}
	return cat.PostAPI(chapterInfoApiPoint, map[string]string{"chapter_id": chapterId, "chapter_command": command})
}

func (cat *Ciweimao) AutoRegV2Api(android string) (gjson.Result, error) {
	return cat.PostAPI(autoRegV2ApiPoint, map[string]string{"gender": "1", "channel": "oppo", "uuid": "android " + android})
}

func (cat *Ciweimao) BookShelfIdListApi() (gjson.Result, error) {
	return cat.PostAPI(bookshelfListApiPoint, nil)
}
func (cat *Ciweimao) BookShelfListApi(shelfId string) (gjson.Result, error) {
	return cat.PostAPI(bookshelfBookListApiPoint, map[string]string{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev"})
}

func (cat *Ciweimao) UseGeetestInfoApi(loginName string) (gjson.Result, error) {
	return cat.PostAPI(useGeetestApiPoint, map[string]string{"login_name": loginName})
}
func (cat *Ciweimao) BookmarkListApi(bookID string, page string) (gjson.Result, error) {
	return cat.PostAPI("/book/get_bookmark_list", map[string]string{"count": "10", "book_id": bookID, "page": page})
}
func (cat *Ciweimao) TsukkomiNumApi(chapterID string) (gjson.Result, error) {
	return cat.PostAPI("/chapter/get_tsukkomi_num", map[string]string{"chapter_id": chapterID})
}

func (cat *Ciweimao) BdaudioInfoApi(bookID string) (gjson.Result, error) {
	return cat.PostAPI("/reader/bdaudio_info", map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) AddReadbookApi(bookID string, readTimes string, getTime string) (gjson.Result, error) {
	return cat.PostAPI("/reader/add_readbook", map[string]string{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *Ciweimao) SetLastReadChapterApi(lastReadChapterID string, bookID string) (gjson.Result, error) {
	return cat.PostAPI("/bookshelf/set_last_read_chapter", map[string]string{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *Ciweimao) PostPrivacyPolicyVersionApi() (gjson.Result, error) {
	return cat.PostAPI("/setting/privacy_policy_version", map[string]string{"privacy_policy_version": "1"})
}

func (cat *Ciweimao) PostPropInfoApi() (gjson.Result, error) {
	return cat.PostAPI("/reader/get_prop_info", nil)
}

func (cat *Ciweimao) MetaDataApi() (gjson.Result, error) {
	return cat.PostAPI("/meta/get_meta_data", nil)
}

func (cat *Ciweimao) VersionApi() (gjson.Result, error) {
	return cat.PostAPI("/setting/get_version", nil)
}

func (cat *Ciweimao) StartpageUrlListApi() (gjson.Result, error) {
	return cat.PostAPI("/setting/get_startpage_url_list", nil)
}

func (cat *Ciweimao) ThirdPartySwitchApi() (gjson.Result, error) {
	return cat.PostAPI("/setting/thired_party_switch", nil)
}
