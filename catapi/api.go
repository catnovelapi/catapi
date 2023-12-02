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
		return response.GetByte(), nil
	}
	return nil, fmt.Errorf("download cover error:%s\nurl:%s\n", "retry 5 times", url)
}
func (cat *Ciweimao) AccountInfoApi() (gjson.Result, error) {
	if accountInfo, err := cat.Req.Post(accountInfoApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else {
		return accountInfo.Get("data.reader_info"), nil
	}
}

// Deprecated: use ChaptersCatalogV2Api instead
func (cat *Ciweimao) ChaptersCatalogApi(bookId string) (gjson.Result, error) {
	return cat.Req.Post(catalogApiPoint, map[string]any{"book_id": bookId})
}

func (cat *Ciweimao) ChaptersCatalogV2Api(bookId string) (gjson.Result, error) {
	if catalog, err := cat.Req.Post(catalogNewApiPoint, map[string]any{"book_id": bookId}); err != nil {
		return gjson.Result{}, err
	} else {
		return catalog.Get("data.chapter_list"), nil
	}
}

func (cat *Ciweimao) BookInfoApiByBookId(bookId string) (gjson.Result, error) {
	query := map[string]any{"use_daguan": "0", "module_id": "20005", "tab_type": "200", "recommend": "module_list", "carousel_position": "", "book_id": bookId}
	if len(bookId) != 9 {
		return gjson.Result{}, fmt.Errorf("bookId length is not 9")
	} else if bookInfo, err := cat.Req.Post(bookInfoApiPoint, query); err != nil {
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
	if result, err := cat.Req.Post(reviewListApiPoint, map[string]any{"book_id": bookId, "count": "10", "page": page, "type": "1"}); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.review_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review list is empty")
	} else {
		return result.Get("data.review_list"), nil
	}
}
func (cat *Ciweimao) ReviewCommentListApi(reviewId string, page string) (gjson.Result, error) {
	if result, err := cat.Req.Post(bookReviewCommentListApiPoint, map[string]any{"review_id": reviewId, "count": "10", "page": page}); err != nil {
		return gjson.Result{}, err
	} else {
		return result.Get("data"), nil
	}
}
func (cat *Ciweimao) ReviewCommentReplyListApi(commentId string, page string) (gjson.Result, error) {
	if result, err := cat.Req.Post(reviewCommentReplyListApiPoint, map[string]any{"comment_id": commentId, "count": "10", "page": page}); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.review_comment_reply_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review comment reply list is empty")
	} else {
		return result.Get("data.review_comment_reply_list"), nil
	}
}
func (cat *Ciweimao) SearchByKeywordApi(keyword, page string) (gjson.Result, error) {
	query := map[string]any{"count": "10", "page": page, "category_index": "0", "key": keyword}
	if search, err := cat.Req.Post(searchBookApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(search.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book is empty")
	} else {
		return search.Get("data.book_list"), nil
	}
}

func (cat *Ciweimao) RedTagBookListApi(tagName, page string) (gjson.Result, error) {
	query := map[string]any{"count": "10", "page": page, "type": "0", "tag": tagName}
	if result, err := cat.Req.Post(redTagApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book red tag is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}
func (cat *Ciweimao) YellowAndBlueTagBookListApi(tagName, filterWord, page string) (gjson.Result, error) {
	query := map[string]any{"filter_word": filterWord, "count": "10", "use_daguan": "0", "page": page,
		"is_paid": "", "category_index": "0", "key": "", "filter_uptime": "", "up_status": "", "order": ""}
	query["tags"] = `[{"filter":"1","tag":"` + tagName + `"}]`
	if result, err := cat.Req.Post(searchBookApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book yellow tag is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}
func (cat *Ciweimao) SignupApi(account string, password string) (gjson.Result, error) {
	return cat.Req.Post(loginApiPoint, map[string]any{"login_name": account, "passwd": password})
}
func (cat *Ciweimao) ChapterCommandApi(chapterId string) (string, error) {
	if commandInfo, err := cat.Req.Post(chapterCommandApiPoint, map[string]any{"chapter_id": chapterId}); err != nil {
		return "", fmt.Errorf("ChapterID:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	} else {
		return commandInfo.Get("data.command").String(), nil
	}
}

func (cat *Ciweimao) TsukkomiNumApi(chapterID string) (gjson.Result, error) {
	return cat.Req.Post(chapterTsukkomiNumApiPoint, map[string]any{"chapter_id": chapterID})
}

func (cat *Ciweimao) contentInfoApi(chapterId string) (gjson.Result, string, error) {
	command, err := cat.ChapterCommandApi(chapterId)
	if err != nil {
		return gjson.Result{}, "", fmt.Errorf("ChapterTitle:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	}
	chapterInfo, err := cat.Req.Post(chapterInfoApiPoint, map[string]any{"chapter_id": chapterId, "chapter_command": command})
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
	query := map[string]any{"gender": "1", "channel": "oppo", "uuid": "android " + android}
	if autoReg, err := cat.Req.Post(autoRegV2ApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else {
		return autoReg.Get("data"), nil
	}
}

func (cat *Ciweimao) BookShelfIdListApi() (gjson.Result, error) {
	if bookshelf, err := cat.Req.Post(bookshelfListApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else if len(bookshelf.Get("data.shelf_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return bookshelf.Get("data.shelf_list"), nil
	}
}
func (cat *Ciweimao) BookShelfListApi(shelfId string) (gjson.Result, error) {
	query := map[string]any{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev", "order": "last_read_time", "count": "999", "page": "0"}
	if result, err := cat.Req.Post(bookshelfBookListApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}

func (cat *Ciweimao) GetVersionApi() (string, error) {
	if result, err := cat.Req.Post(getVersionApiPoint, nil); err != nil {
		return "", err
	} else {
		return result.Get("data.android_version").String(), nil
	}
}
func (cat *Ciweimao) CheckVersionApi() (gjson.Result, error) {
	if result, err := cat.Req.Post(checkVersionApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else {
		return result.Get("data"), nil
	}
}

func (cat *Ciweimao) UseGeetestInfoApi(loginName string) (int, error) {
	useGeetest, err := cat.Req.Post(useGeetestApiPoint, map[string]any{"login_name": loginName})
	if err != nil {
		return 0, err
	}
	return int(useGeetest.Get("data.need_use_geetest").Int()), nil
}
func (cat *Ciweimao) BookmarkListApi(bookID string, page string) (gjson.Result, error) {
	return cat.Req.Post("/book/get_bookmark_list", map[string]any{"count": "10", "book_id": bookID, "page": page})
}
func (cat *Ciweimao) BdaudioInfoApi(bookID string) (gjson.Result, error) {
	return cat.Req.Post("/reader/bdaudio_info", map[string]any{"book_id": bookID})
}

func (cat *Ciweimao) AddReadbookApi(bookID string, readTimes string, getTime string) (gjson.Result, error) {
	return cat.Req.Post("/reader/add_readbook", map[string]any{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *Ciweimao) SetLastReadChapterApi(lastReadChapterID string, bookID string) (gjson.Result, error) {
	return cat.Req.Post("/bookshelf/set_last_read_chapter", map[string]any{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *Ciweimao) PostPrivacyPolicyVersionApi() (gjson.Result, error) {
	return cat.Req.Post("/setting/privacy_policy_version", map[string]any{"privacy_policy_version": "1"})
}

func (cat *Ciweimao) PostPropInfoApi() (gjson.Result, error) {
	return cat.Req.Post("/reader/get_prop_info", nil)
}

func (cat *Ciweimao) MetaDataApi() (gjson.Result, error) {
	return cat.Req.Post("/meta/get_meta_data", nil)
}
func (cat *Ciweimao) StartpageUrlListApi() (gjson.Result, error) {
	return cat.Req.Post("/setting/get_startpage_url_list", nil)
}

func (cat *Ciweimao) ThirdPartySwitchApi() (gjson.Result, error) {
	return cat.Req.Post("/setting/thired_party_switch", nil)
}
