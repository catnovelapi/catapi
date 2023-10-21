package catapi

import (
	"github.com/catnovelapi/catapi/options"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

func (cat *Ciweimao) AccountInfoApi() gjson.Result {
	return cat.post(accountInfoApiPoint, nil)
}

func (cat *Ciweimao) CatalogByBookIdApi(bookID string) gjson.Result {
	return cat.post(catalogApiPoint, map[string]any{"book_id": bookID})
}

func (cat *Ciweimao) CatalogByBookIdNewApi(bookID string) gjson.Result {
	return cat.post(catalogNewApiPoint, map[string]any{"book_id": bookID})
}

func (cat *Ciweimao) BookInfoApi(bookId string) gjson.Result {
	return cat.post(bookInfoApiPoint, map[string]any{"book_id": bookId})
}

func (cat *Ciweimao) SearchByKeywordApi(keyword string, page, categoryIndex int) gjson.Result {
	return cat.post(searchBookApiPoint, map[string]any{"count": "10", "page": page, "category_index": categoryIndex, "key": keyword})
}

func (cat *Ciweimao) SignupApi(account, password string) gjson.Result {
	return cat.post(loginApiPoint, map[string]any{"login_name": account, "passwd": password})
}

func (cat *Ciweimao) ChapterCommandApi(chapterId string) gjson.Result {
	return cat.post(chapterCommandApiPoint, map[string]any{"chapter_id": chapterId})
}

func (cat *Ciweimao) ChapterInfoApi(chapterId string, command string) gjson.Result {
	return cat.post(chapterInfoApiPoint, map[string]any{"chapter_id": chapterId, "chapter_command": command})
}

func (cat *Ciweimao) AutoRegV2Api() gjson.Result {
	return cat.post(autoRegV2ApiPoint, map[string]any{"gender": "1", "channel": "oppo", "uuid": "android " + uuid.New().String()})
}

func (cat *Ciweimao) BookShelfIdListApi() gjson.Result {
	return cat.post(bookshelfListApiPoint, nil)
}
func (cat *Ciweimao) BookShelfListApi(shelfId string) gjson.Result {
	return cat.post(bookshelfBookListApiPoint, map[string]any{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev"})
}

func (cat *Ciweimao) UseGeetestInfoApi(loginName string) gjson.Result {
	return cat.post(useGeetestApiPoint, map[string]any{"login_name": loginName}, options.NoDecode())
}
func (cat *Ciweimao) BookmarkListApi(bookID string, page string) gjson.Result {
	return cat.post("/book/get_bookmark_list", map[string]any{"count": "10", "book_id": bookID, "page": page})
}
func (cat *Ciweimao) TsukkomiNumApi(chapterID string) gjson.Result {
	return cat.post("/chapter/get_tsukkomi_num", map[string]any{"chapter_id": chapterID})
}

func (cat *Ciweimao) BdaudioInfoApi(bookID string) gjson.Result {
	return cat.post("/reader/bdaudio_info", map[string]any{"book_id": bookID})
}

func (cat *Ciweimao) AddReadbookApi(bookID string, readTimes string, getTime string) gjson.Result {
	return cat.post("/reader/add_readbook", map[string]any{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *Ciweimao) SetLastReadChapterApi(lastReadChapterID string, bookID string) gjson.Result {
	return cat.post("/bookshelf/set_last_read_chapter", map[string]any{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *Ciweimao) PostPrivacyPolicyVersionApi() gjson.Result {
	return cat.post("/setting/privacy_policy_version", map[string]any{"privacy_policy_version": "1"})
}

func (cat *Ciweimao) PostPropInfoApi() gjson.Result {
	return cat.post("/reader/get_prop_info", nil)
}

func (cat *Ciweimao) MetaDataApi() gjson.Result {
	return cat.post("/meta/get_meta_data", nil)
}

func (cat *Ciweimao) VersionApi() gjson.Result {
	return cat.post("/setting/get_version", nil)
}

func (cat *Ciweimao) StartpageUrlListApi() gjson.Result {
	return cat.post("/setting/get_startpage_url_list", nil)
}

func (cat *Ciweimao) ThirdPartySwitchApi() gjson.Result {
	return cat.post("/setting/thired_party_switch", nil)
}
