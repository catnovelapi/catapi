package catapi

import (
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type Ciweimao struct {
	Debug         bool
	Proxy         string
	Host          string
	Version       string
	LoginToken    string
	Account       string
	BuilderClient *resty.Client
}

func (cat *Ciweimao) AccountInfoApi() (gjson.Result, error) {
	return cat.PostAPI(accountInfoApiPoint, nil)
}

func (cat *Ciweimao) CatalogByBookIdApi(bookID string) (gjson.Result, error) {
	return cat.PostAPI(catalogApiPoint, map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) CatalogByBookIdNewApi(bookID string) (gjson.Result, error) {
	return cat.PostAPI(catalogNewApiPoint, map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) BookInfoApi(bookId string) (gjson.Result, error) {
	return cat.PostAPI(bookInfoApiPoint, map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) SearchByKeywordApi(keyword, page, categoryIndex string) (gjson.Result, error) {
	return cat.PostAPI(searchBookApiPoint, map[string]string{"count": "10", "page": page, "category_index": categoryIndex, "key": keyword})
}

func (cat *Ciweimao) SignupApi(account, password string) (gjson.Result, error) {
	return cat.PostAPI(loginApiPoint, map[string]string{"login_name": account, "passwd": password})
}

func (cat *Ciweimao) ChapterCommandApi(chapterId string) (gjson.Result, error) {
	return cat.PostAPI(chapterCommandApiPoint, map[string]string{"chapter_id": chapterId})
}

func (cat *Ciweimao) ChapterInfoApi(chapterId string, command string) (gjson.Result, error) {
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

func (cat *Ciweimao) UseGeetestInfoApi(loginName string) (*resty.Response, error) {
	return cat.Post(useGeetestApiPoint, map[string]string{"login_name": loginName})
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
