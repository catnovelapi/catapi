package catapi

import (
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
)

type Ciweimao struct {
	Req *CiweimaoRequest
}

func (cat *Ciweimao) DownloadCover(url string) ([]byte, error) {
	for i := 0; i < 5; i++ {
		response, err := cat.Req.BuilderClient.R().Get(url)
		if err != nil {
			fmt.Printf("download cover error,retry %d times\n", i)
			continue
		}
		return response.Body(), nil
	}
	return nil, fmt.Errorf("download cover error:%s\nurl:%s\n", "retry 5 times", url)
}
func (cat *Ciweimao) AccountInfoApi() (gjson.Result, error) {
	if accountInfo, err := cat.Req.PostAPI(accountInfoApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else {
		return accountInfo.Get("data.reader_info"), nil
	}
}

// Deprecated: use ChaptersCatalogV2Api instead
func (cat *Ciweimao) ChaptersCatalogApi(bookId string) (gjson.Result, error) {
	return cat.Req.PostAPI(catalogApiPoint, map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) ChaptersCatalogV2Api(bookId string) (gjson.Result, error) {
	if catalog, err := cat.Req.PostAPI(catalogNewApiPoint, map[string]string{"book_id": bookId}); err != nil {
		return gjson.Result{}, err
	} else {
		return catalog.Get("data.chapter_list"), nil
	}
}

func (cat *Ciweimao) BookInfoApiByBookId(bookId string) (gjson.Result, error) {
	query := map[string]string{"use_daguan": "0", "module_id": "20005", "tab_type": "200", "recommend": "module_list", "carousel_position": "", "book_id": bookId}
	if len(bookId) != 9 {
		return gjson.Result{}, fmt.Errorf("bookId length is not 9")
	} else if bookInfo, err := cat.Req.PostAPI(bookInfoApiPoint, query); err != nil {
		return gjson.Result{}, fmt.Errorf("bookId:%s,获取书籍信息失败:%s", bookId, err.Error())
	} else {
		return bookInfo.Get("data.book_info"), nil
	}
}

func (cat *Ciweimao) BookInfoApiByBookURL(url string) (gjson.Result, error) {
	if bi := regexp.MustCompile(`(\d{9})`).FindStringSubmatch(url); len(bi) < 2 {
		return gjson.Result{}, fmt.Errorf("bookId is empty")
	} else {
		return cat.BookInfoApiByBookId(bi[1])
	}
}

func (cat *Ciweimao) ReviewListApi(bookId string, page string) (gjson.Result, error) {
	return cat.Req.PostAPI(reviewListApiPoint, map[string]string{"book_id": bookId, "count": "10", "page": page, "type": "1"})
}
func (cat *Ciweimao) ReviewCommentListApi(reviewId string, page string) (gjson.Result, error) {
	return cat.Req.PostAPI(bookReviewCommentListApiPoint, map[string]string{"review_id": reviewId, "count": "10", "page": page})
}
func (cat *Ciweimao) ReviewCommentReplyListApi(commentId string, page string) (gjson.Result, error) {
	return cat.Req.PostAPI(reviewCommentReplyListApiPoint, map[string]string{"comment_id": commentId, "count": "10", "page": page})
}
func (cat *Ciweimao) SearchByKeywordApi(keyword, page string) (gjson.Result, error) {
	query := map[string]string{"count": "10", "page": page, "category_index": "0", "key": keyword}
	if search, err := cat.Req.PostAPI(searchBookApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(search.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book is empty")
	} else {
		return search.Get("data.book_list"), nil
	}
}

func (cat *Ciweimao) SearchByTagApi(tagName, page string) (gjson.Result, error) {
	return cat.Req.PostAPI(searchBookTagApiPoint, map[string]string{"count": "10", "page": page, "type": "0", "tag": tagName})
}
func (cat *Ciweimao) SignupApi(account string, password string) (gjson.Result, error) {
	return cat.Req.PostAPI(loginApiPoint, map[string]string{"login_name": account, "passwd": password})
}
func (cat *Ciweimao) ChapterCommandApi(chapterId string) (string, error) {
	if commandInfo, err := cat.Req.PostAPI(chapterCommandApiPoint, map[string]string{"chapter_id": chapterId}); err != nil {
		return "", fmt.Errorf("ChapterID:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	} else {
		return commandInfo.Get("data.command").String(), nil
	}
}

func (cat *Ciweimao) contentInfoApi(chapterId string) (gjson.Result, string, error) {
	command, err := cat.ChapterCommandApi(chapterId)
	if err != nil {
		return gjson.Result{}, "", fmt.Errorf("ChapterTitle:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	}
	chapterInfo, err := cat.Req.PostAPI(chapterInfoApiPoint, map[string]string{"chapter_id": chapterId, "chapter_command": command})
	if err != nil {
		return gjson.Result{}, "", err
	}
	return chapterInfo.Get("data.chapter_info"), command, nil
}

func (cat *Ciweimao) ChapterContentApi(chapterId string) (string, error) {
	chapterInfo, command, err := cat.contentInfoApi(chapterId)
	if err != nil {
		return "", err
	}
	chapterInfoText, err := cat.Req.DecodeEncryptText(chapterInfo.Get("txt_content").String(), command)
	if err != nil {
		return "", err
	}
	return chapterInfoText, nil
}
func (cat *Ciweimao) ChapterInfoApi(chapterId string) (gjson.Result, error) {
	chapterInfo, _, err := cat.contentInfoApi(chapterId)
	return chapterInfo, err
}

func (cat *Ciweimao) AutoRegV2Api(android string) (gjson.Result, error) {
	return cat.Req.PostAPI(autoRegV2ApiPoint, map[string]string{"gender": "1", "channel": "oppo", "uuid": "android " + android})
}

func (cat *Ciweimao) BookShelfIdListApi() (gjson.Result, error) {
	if bookshelf, err := cat.Req.PostAPI(bookshelfListApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else if len(bookshelf.Get("data.shelf_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return bookshelf.Get("data.shelf_list"), nil
	}
}
func (cat *Ciweimao) BookShelfListApi(shelfId string) (gjson.Result, error) {
	query := map[string]string{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev", "order": "last_read_time", "count": "999", "page": "0"}
	if result, err := cat.Req.PostAPI(bookshelfBookListApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}

func (cat *Ciweimao) UseGeetestInfoApi(loginName string) (int, error) {
	useGeetest, err := cat.Req.PostAPI(useGeetestApiPoint, map[string]string{"login_name": loginName})
	if err != nil {
		return 0, err
	}
	return int(useGeetest.Get("data.need_use_geetest").Int()), nil
}
func (cat *Ciweimao) BookmarkListApi(bookID string, page string) (gjson.Result, error) {
	return cat.Req.PostAPI("/book/get_bookmark_list", map[string]string{"count": "10", "book_id": bookID, "page": page})
}
func (cat *Ciweimao) TsukkomiNumApi(chapterID string) (gjson.Result, error) {
	return cat.Req.PostAPI("/chapter/get_tsukkomi_num", map[string]string{"chapter_id": chapterID})
}

func (cat *Ciweimao) BdaudioInfoApi(bookID string) (gjson.Result, error) {
	return cat.Req.PostAPI("/reader/bdaudio_info", map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) AddReadbookApi(bookID string, readTimes string, getTime string) (gjson.Result, error) {
	return cat.Req.PostAPI("/reader/add_readbook", map[string]string{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *Ciweimao) SetLastReadChapterApi(lastReadChapterID string, bookID string) (gjson.Result, error) {
	return cat.Req.PostAPI("/bookshelf/set_last_read_chapter", map[string]string{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *Ciweimao) PostPrivacyPolicyVersionApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/setting/privacy_policy_version", map[string]string{"privacy_policy_version": "1"})
}

func (cat *Ciweimao) PostPropInfoApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/reader/get_prop_info", nil)
}

func (cat *Ciweimao) MetaDataApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/meta/get_meta_data", nil)
}

func (cat *Ciweimao) VersionApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/setting/get_version", nil)
}

func (cat *Ciweimao) StartpageUrlListApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/setting/get_startpage_url_list", nil)
}

func (cat *Ciweimao) ThirdPartySwitchApi() (gjson.Result, error) {
	return cat.Req.PostAPI("/setting/thired_party_switch", nil)
}
