package catapi

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
)

// post 发送post请求
func (cat *API) post(url string, data map[string]string) (gjson.Result, error) {
	req := cat.builderClient.R()
	if data != nil {
		req.SetQueryParams(data)
	}
	response, err := req.Post(url)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("request error: %s", err.Error())
	}

	if result := gjson.Parse(response.String()); result.Get("code").String() != "100000" {
		return result, fmt.Errorf("response error: %s", result.Get("tip").String())
	} else {
		return result, nil
	}
}
func (cat *API) DownloadCover(url string) ([]byte, error) {
	for i := 0; i < 5; i++ {
		response, err := cat.builderClient.R().Get(url)
		if err != nil {
			fmt.Printf("download cover error,retry %d times\n", i)
			continue
		}
		return response.GetByte(), nil
	}
	return nil, fmt.Errorf("download cover error:%s\nurl:%s\n", "retry 5 times", url)
}

// AccountInfoApi 获取账号信息
func (cat *API) AccountInfoApi() (gjson.Result, error) {
	if accountInfo, err := cat.post(accountInfoApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else {
		return accountInfo.Get("data.reader_info"), nil
	}
}

// Deprecated: use ChaptersCatalogV2Api instead
func (cat *API) ChaptersCatalogApi(bookId string) (gjson.Result, error) {
	return cat.post(catalogApiPoint, map[string]string{"book_id": bookId})
}

// ChaptersCatalogV2Api 获取章节列表,需要传入书籍ID
func (cat *API) ChaptersCatalogV2Api(bookId string) (gjson.Result, error) {
	if catalog, err := cat.post(catalogNewApiPoint, map[string]string{"book_id": bookId}); err != nil {
		return gjson.Result{}, err
	} else {
		return catalog.Get("data.chapter_list"), nil
	}
}

// BookInfoApiByBookId 通过书籍ID获取书籍信息
func (cat *API) BookInfoApiByBookId(bookId string) (gjson.Result, error) {
	query := map[string]string{"use_daguan": "0", "module_id": "20005", "tab_type": "200", "recommend": "module_list", "carousel_position": "", "book_id": bookId}
	if len(bookId) != 9 {
		return gjson.Result{}, fmt.Errorf("bookId length is not 9")
	} else if bookInfo, err := cat.post(bookInfoApiPoint, query); err != nil {
		return gjson.Result{}, fmt.Errorf("bookId:%s,获取书籍信息失败:%s", bookId, err.Error())
	} else {
		return bookInfo.Get("data.book_info"), nil
	}
}

// BookInfoApiByBookURL 通过书籍URL获取书籍信息
func (cat *API) BookInfoApiByBookURL(url string) (gjson.Result, error) {
	if bi := regexp.MustCompile(`(\d{9})`).FindStringSubmatch(url); len(bi) < 2 {
		return gjson.Result{}, fmt.Errorf("bookId is empty")
	} else {
		return cat.BookInfoApiByBookId(bi[1])
	}
}

// ReviewListApi 获取书评列表,需要传入书籍ID和页码
func (cat *API) ReviewListApi(bookId string, page string) (gjson.Result, error) {
	if result, err := cat.post(reviewListApiPoint, map[string]string{"book_id": bookId, "count": "10", "page": page, "type": "1"}); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.review_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review list is empty")
	} else {
		return result.Get("data.review_list"), nil
	}
}

// ReviewCommentListApi 获取书评评论列表,需要传入书评ID和页码
func (cat *API) ReviewCommentListApi(reviewId string, page string) (gjson.Result, error) {
	if result, err := cat.post(bookReviewCommentListApiPoint, map[string]string{"review_id": reviewId, "count": "10", "page": page}); err != nil {
		return gjson.Result{}, err
	} else {
		return result.Get("data"), nil
	}
}

// ReviewCommentReplyListApi 获取书评回复列表,需要传入书评ID和页码
func (cat *API) ReviewCommentReplyListApi(commentId string, page string) (gjson.Result, error) {
	if result, err := cat.post(reviewCommentReplyListApiPoint, map[string]string{"comment_id": commentId, "count": "10", "page": page}); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.review_comment_reply_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review comment reply list is empty")
	} else {
		return result.Get("data.review_comment_reply_list"), nil
	}
}

// SearchByKeywordApi 搜索书籍,需要传入关键字和页码
func (cat *API) SearchByKeywordApi(keyword, page string) (gjson.Result, error) {
	query := map[string]string{"count": "10", "page": page, "category_index": "0", "key": keyword}
	if search, err := cat.post(searchBookApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(search.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book is empty")
	} else {
		return search.Get("data.book_list"), nil
	}
}

// RedTagBookListApi 获取红标签书籍列表
func (cat *API) RedTagBookListApi(tagName, page string) (gjson.Result, error) {
	query := map[string]string{"count": "10", "page": page, "type": "0", "tag": tagName}
	if result, err := cat.post(redTagApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book red tag is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}

// YellowAndBlueTagBookListApi 获取黄蓝标签书籍列表
func (cat *API) YellowAndBlueTagBookListApi(tagName, filterWord, page string) (gjson.Result, error) {
	query := map[string]string{"filter_word": filterWord, "count": "10", "use_daguan": "0", "page": page,
		"is_paid": "", "category_index": "0", "key": "", "filter_uptime": "", "up_status": "", "order": ""}
	query["tags"] = `[{"filter":"1","tag":"` + tagName + `"}]`
	if result, err := cat.post(searchBookApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("search book yellow tag is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}

// SignupApi 注册账号,需要传入账号和密码
func (cat *API) SignupApi(account string, password string) (gjson.Result, error) {
	return cat.post(loginApiPoint, map[string]string{"login_name": account, "passwd": password})
}

// ChapterCommandApi 获取章节command,需要传入章节ID
func (cat *API) ChapterCommandApi(chapterId string) (gjson.Result, error) {
	if commandInfo, err := cat.post(chapterCommandApiPoint, map[string]string{"chapter_id": chapterId}); err != nil {
		return gjson.Result{}, fmt.Errorf("ChapterID:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	} else {
		return commandInfo.Get("data"), nil
	}
}

// ContentInfoApi 获取章节内容,需要传入章节ID
func (cat *API) ContentInfoApi(chapterId string) (*ContentInfoTemplate, error) {
	command, err := cat.ChapterCommandApi(chapterId)
	if err != nil {
		return nil, fmt.Errorf("ChapterTitle:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	}
	return cat.ContentInfoByCommandApi(chapterId, command.Get("chapter_command").String())
}

func (cat *API) ContentInfoByCommandApi(chapterId, command string) (*ContentInfoTemplate, error) {
	params := map[string]string{"chapter_id": chapterId, "chapter_command": command}
	chapter, err := cat.post(chapterInfoApiPoint, params)
	if err != nil {
		return nil, err
	}
	var chapterInfo *ContentInfoTemplate
	err = json.Unmarshal([]byte(chapter.Get("data.chapter_info").String()), &chapterInfo)
	if err != nil {
		return nil, err
	}
	if chapterInfo.TxtContent == "" {
		return nil, fmt.Errorf("ChapterTitle:%s,获取章节内容失败,tips:%s", chapterId, "txt_content is empty")
	}
	chapterInfoText, ok := DecodeEncryptText(chapterInfo.TxtContent, command)
	if ok != nil {
		return nil, ok
	} else {
		chapterInfo.TxtContent = chapterInfoText
	}
	return chapterInfo, nil
}

func (cat *API) TsukkomiNumApi(chapterID string) (gjson.Result, error) {
	return cat.post(chapterTsukkomiNumApiPoint, map[string]string{"chapter_id": chapterID})
}

func (cat *API) RegV2Api() (gjson.Result, error) {
	query := map[string]string{"gender": "1", "channel": "oppo", "uuid": "android " + cat.Client.AndroidID()}
	if autoReg, err := cat.post(autoRegV2ApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else {
		return autoReg.Get("data"), nil
	}
}

func (cat *API) BookShelfIdListApi() (gjson.Result, error) {
	if bookshelf, err := cat.post(bookshelfListApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else if len(bookshelf.Get("data.shelf_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return bookshelf.Get("data.shelf_list"), nil
	}
}
func (cat *API) BookShelfListApi(shelfId string) (gjson.Result, error) {
	query := map[string]string{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev", "order": "last_read_time", "count": "999", "page": "0"}
	if result, err := cat.post(bookshelfBookListApiPoint, query); err != nil {
		return gjson.Result{}, err
	} else if len(result.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("bookshelf is empty")
	} else {
		return result.Get("data.book_list"), nil
	}
}

func (cat *API) GetVersionApi() (string, error) {
	if result, err := cat.post(getVersionApiPoint, nil); err != nil {
		return "", err
	} else {
		return result.Get("data.android_version").String(), nil
	}
}
func (cat *API) CheckVersionApi() (gjson.Result, error) {
	if result, err := cat.post(checkVersionApiPoint, nil); err != nil {
		return gjson.Result{}, err
	} else {
		return result.Get("data"), nil
	}
}

func (cat *API) UseGeetestInfoApi(loginName string) (int, error) {
	useGeetest, err := cat.post(useGeetestApiPoint, map[string]string{"login_name": loginName})
	if err != nil {
		return 0, err
	}
	return int(useGeetest.Get("data.need_use_geetest").Int()), nil
}
func (cat *API) BookmarkListApi(bookID string, page string) (gjson.Result, error) {
	return cat.post("/book/get_bookmark_list", map[string]string{"count": "10", "book_id": bookID, "page": page})
}
func (cat *API) BdaudioInfoApi(bookID string) (gjson.Result, error) {
	return cat.post("/reader/bdaudio_info", map[string]string{"book_id": bookID})
}

func (cat *API) AddReadbookApi(bookID string, readTimes string, getTime string) (gjson.Result, error) {
	return cat.post("/reader/add_readbook", map[string]string{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *API) SetLastReadChapterApi(lastReadChapterID string, bookID string) (gjson.Result, error) {
	return cat.post("/bookshelf/set_last_read_chapter", map[string]string{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *API) PostPrivacyPolicyVersionApi() (gjson.Result, error) {
	return cat.post("/setting/privacy_policy_version", map[string]string{"privacy_policy_version": "1"})
}

func (cat *API) PostPropInfoApi() (gjson.Result, error) {
	return cat.post("/reader/get_prop_info", nil)
}

func (cat *API) MetaDataApi() (gjson.Result, error) {
	return cat.post("/meta/get_meta_data", nil)
}
func (cat *API) StartpageUrlListApi() (gjson.Result, error) {
	return cat.post("/setting/get_startpage_url_list", nil)
}

func (cat *API) ThirdPartySwitchApi() (gjson.Result, error) {
	return cat.post("/setting/thired_party_switch", nil)
}
