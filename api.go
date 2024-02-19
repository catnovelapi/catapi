package catapi

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
)

// post 发送post请求
func (cat *API) post(url string, data any) (gjson.Result, error) {
	req := cat.builderClient.R()
	if data != nil {
		switch params := data.(type) {
		case map[string]any:
			req.SetQueryParams(params)
		default:
			req.SetBody(data)
		}
	}
	if url == autoRegV2ApiPoint {
		// 暂时删掉login_token和account
		req.SetQueryParam("login_token", "")
		req.SetQueryParam("account", "")
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
func (cat *API) checkbookList(response gjson.Result, err error) (gjson.Result, error) {
	if err != nil {
		return gjson.Result{}, err
	} else if len(response.Get("data.book_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("get book_list is empty")
	} else {
		return response.Get("data.book_list"), nil
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
	return cat.post(catalogApiPoint, map[string]any{"book_id": bookId})
}

// ChaptersCatalogV2Api 获取章节列表,需要传入书籍ID
func (cat *API) ChaptersCatalogV2Api(bookId string) (*ChapterListTemplate, error) {
	catalog, err := cat.post(catalogNewApiPoint, map[string]any{"book_id": bookId})
	if err != nil {
		return nil, err
	} else {
		var chapterList *ChapterListTemplate
		err = json.Unmarshal([]byte(catalog.String()), &chapterList)
		if err != nil {
			return nil, err
		}
		if chapterList.Code != "100000" {
			return nil, fmt.Errorf("bookId:%s,获取章节列表失败:%v", bookId, chapterList.Tip)
		} else if len(chapterList.Data.ChapterList) == 0 {
			return nil, fmt.Errorf("bookId:%s,获取章节列表失败:%v", bookId, "chapter_list is empty")
		}
		return chapterList, nil
	}
}

// BookInfoApi 通过书籍ID获取书籍信息
func (cat *API) BookInfoApi(bookId string) (*BookInfoTemplate, error) {
	params := BookInfoQuery{BookId: bookId, ModuleId: "20005", TabType: "200", Recommend: "module_list", UseDaguan: "0"}
	bookInfo, err := cat.post(bookInfoApiPoint, params)
	if err != nil {
		return nil, fmt.Errorf("bookId:%s,获取书籍信息失败:%s", bookId, err.Error())
	}
	var bookInfoTemplate *BookInfoTemplate
	err = json.Unmarshal([]byte(bookInfo.Get("data.book_info").String()), &bookInfoTemplate)
	if err != nil {
		return nil, err
	}
	if bookInfoTemplate.BookName == "" {
		return nil, fmt.Errorf("bookId:%s,获取书籍信息失败:%s", bookId, "book_name is empty")
	}
	return bookInfoTemplate, nil
}

// BookInfoApiByBookURL 通过书籍URL获取书籍信息
func (cat *API) BookInfoApiByBookURL(url string) (*BookInfoTemplate, error) {
	if bi := regexp.MustCompile(`(\d{9})`).FindStringSubmatch(url); len(bi) < 2 {
		return nil, fmt.Errorf("bookId is empty")
	} else {
		return cat.BookInfoApi(bi[1])
	}
}

// ReviewListApi 获取书评列表,需要传入书籍ID和页码
func (cat *API) ReviewListApi(bookId string, page string) (gjson.Result, error) {
	response, err := cat.post(reviewListApiPoint, map[string]any{"book_id": bookId, "count": "10", "page": page, "type": "1"})
	if err != nil {
		return gjson.Result{}, err
	} else if len(response.Get("data.review_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review list is empty")
	} else {
		return response.Get("data.review_list"), nil
	}
}

// ReviewCommentListApi 获取书评评论列表,需要传入书评ID和页码
func (cat *API) ReviewCommentListApi(reviewId string, page string) (gjson.Result, error) {
	if result, err := cat.post(bookReviewCommentListApiPoint, map[string]any{"review_id": reviewId, "count": "10", "page": page}); err != nil {
		return gjson.Result{}, err
	} else {
		return result.Get("data"), nil
	}
}

// ReviewCommentReplyListApi 获取书评回复列表,需要传入书评ID和页码
func (cat *API) ReviewCommentReplyListApi(commentId string, page string) (gjson.Result, error) {
	response, err := cat.post(reviewCommentReplyListApiPoint, map[string]any{"comment_id": commentId, "count": "10", "page": page})
	if err != nil {
		return gjson.Result{}, err
	} else if len(response.Get("data.review_comment_reply_list").Array()) == 0 {
		return gjson.Result{}, fmt.Errorf("review comment reply list is empty")
	} else {
		return response.Get("data.review_comment_reply_list"), nil
	}
}

// SearchByKeywordApi 搜索书籍,需要传入关键字和页码
func (cat *API) SearchByKeywordApi(keyword string, page string) (gjson.Result, error) {
	return cat.checkbookList(cat.post(searchBookApiPoint, SearchKeywordQuery{Count: "10", Page: page, CategoryIndex: "0", Key: keyword}))
}

// RedTagBookListApi 获取红标签书籍列表
func (cat *API) RedTagBookListApi(tagName, page string) (gjson.Result, error) {
	return cat.checkbookList(cat.post(redTagApiPoint, map[string]any{"count": "10", "page": page, "type": "0", "tag": tagName}))
}

// YellowAndBlueTagBookListApi 获取黄蓝标签书籍列表
func (cat *API) YellowAndBlueTagBookListApi(tagName, filterWord, page string) (gjson.Result, error) {
	params := SearchTagsQuery{
		FilterWord:    filterWord,
		Count:         "10",
		Page:          page,
		UseDaguan:     "0",
		IsPaid:        "",
		CategoryIndex: "0",
		Key:           "",
		FilterUptime:  "",
		UpStatus:      "",
		Order:         "",
		Tags:          []SearchTagFilterQuery{{Filter: "1", Tag: tagName}},
	}
	return cat.checkbookList(cat.post(searchBookApiPoint, params))
}

// SignupApi 注册账号,需要传入账号和密码
func (cat *API) SignupApi(account string, password string) (gjson.Result, error) {
	return cat.post(loginApiPoint, map[string]any{"login_name": account, "passwd": password})
}

// ChapterCommandApi 获取章节command,需要传入章节ID
func (cat *API) ChapterCommandApi(chapterId string) (gjson.Result, error) {
	if response, err := cat.post(chapterCommandApiPoint, map[string]any{"chapter_id": chapterId}); err != nil {
		return gjson.Result{}, fmt.Errorf("ChapterID:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	} else {
		return response.Get("data"), nil
	}
}

// ContentInfoApi 获取章节内容,需要传入章节ID
func (cat *API) ContentInfoApi(chapterId string) (*ContentInfoTemplate, error) {
	response, err := cat.ChapterCommandApi(chapterId)
	if err != nil {
		return nil, fmt.Errorf("ChapterTitle:%s,获取章节command失败,tips:%s", chapterId, err.Error())
	}
	return cat.ContentInfoByCommandApi(chapterId, response.Get("command").String())
}

func (cat *API) ContentInfoByCommandApi(chapterId, command string) (*ContentInfoTemplate, error) {
	response, err := cat.post(chapterInfoApiPoint, map[string]any{"chapter_id": chapterId, "chapter_command": command})
	if err != nil {
		return nil, err
	}
	var chapterInfo *ContentInfoTemplate
	err = json.Unmarshal([]byte(response.Get("data.chapter_info").String()), &chapterInfo)
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
	return cat.post(chapterTsukkomiNumApiPoint, map[string]any{"chapter_id": chapterID})
}

func (cat *API) RegV2Api() (gjson.Result, error) {
	query := map[string]any{"gender": "1", "channel": "oppo", "uuid": "android " + cat.Client.AndroidID()}
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
func (cat *API) BookShelfListApi(shelfId, page string) (gjson.Result, error) {
	params := ShelfListQuery{LastModTime: "0", ShelfId: shelfId, Direction: "prev", Order: "last_read_time", Count: "999", Page: page}
	return cat.checkbookList(cat.post(bookshelfBookListApiPoint, params))
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

func (cat *API) UseGeetestInfoApi(loginName string) (gjson.Result, error) {
	useGeetest, err := cat.post(useGeetestApiPoint, map[string]any{"login_name": loginName})
	if err != nil {
		return gjson.Result{}, err
	}
	return useGeetest.Get("data"), nil
}
func (cat *API) BookmarkListApi(bookID string, page string) (gjson.Result, error) {
	return cat.post("/book/get_bookmark_list", map[string]any{"count": "10", "book_id": bookID, "page": page})
}
func (cat *API) BdaudioInfoApi(bookID string) (gjson.Result, error) {
	return cat.post("/reader/bdaudio_info", map[string]any{"book_id": bookID})
}

func (cat *API) AddReadbookApi(bookID string, readTimes string, getTime string) (gjson.Result, error) {
	return cat.post("/reader/add_readbook", map[string]any{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *API) SetLastReadChapterApi(lastReadChapterID string, bookID string) (gjson.Result, error) {
	return cat.post(setLastReadChapterApiPoint, map[string]any{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *API) PostPrivacyPolicyVersionApi() (gjson.Result, error) {
	return cat.post("/setting/privacy_policy_version", map[string]any{"privacy_policy_version": "1"})
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
