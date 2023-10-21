package cattemplate

type AutoRegV2Template struct {
	LoginToken string `json:"login_token"`
	UserCode   string `json:"user_code"`
	ReaderInfo struct {
		ReaderId       string        `json:"reader_id"`
		Account        string        `json:"account"`
		IsBind         string        `json:"is_bind"`
		IsBindQq       string        `json:"is_bind_qq"`
		IsBindWeixin   string        `json:"is_bind_weixin"`
		IsBindHuawei   string        `json:"is_bind_huawei"`
		IsBindApple    string        `json:"is_bind_apple"`
		PhoneNum       string        `json:"phone_num"`
		PhoneCrypto    string        `json:"phone_crypto"`
		MobileVal      string        `json:"mobileVal"`
		Email          string        `json:"email"`
		License        string        `json:"license"`
		ReaderName     string        `json:"reader_name"`
		AvatarUrl      string        `json:"avatar_url"`
		AvatarThumbUrl string        `json:"avatar_thumb_url"`
		BaseStatus     string        `json:"base_status"`
		ExpLv          string        `json:"exp_lv"`
		ExpValue       string        `json:"exp_value"`
		Gender         string        `json:"gender"`
		VipLv          string        `json:"vip_lv"`
		VipValue       string        `json:"vip_value"`
		IsAuthor       string        `json:"is_author"`
		IsUploader     string        `json:"is_uploader"`
		BookAge        string        `json:"book_age"`
		CategoryPrefer []interface{} `json:"category_prefer"`
		UsedDecoration []interface{} `json:"used_decoration"`
		Rank           string        `json:"rank"`
		FirstLoginIp   string        `json:"first_login_ip"`
		Ctime          string        `json:"ctime"`
	} `json:"reader_info"`
	PropInfo struct {
		RestGiftHlb     string `json:"rest_gift_hlb"`
		RestHlb         string `json:"rest_hlb"`
		RestYp          string `json:"rest_yp"`
		RestRecommend   string `json:"rest_recommend"`
		RestTotalBlade  string `json:"rest_total_blade"`
		RestMonthBlade  string `json:"rest_month_blade"`
		RestTotal100    string `json:"rest_total_100"`
		RestTotal588    string `json:"rest_total_588"`
		RestTotal1688   string `json:"rest_total_1688"`
		RestTotal5000   string `json:"rest_total_5000"`
		RestTotal10000  string `json:"rest_total_10000"`
		RestTotal100000 string `json:"rest_total_100000"`
		RestTotal50000  string `json:"rest_total_50000"`
		RestTotal160000 string `json:"rest_total_160000"`
	} `json:"prop_info"`
	IsSetYoung string `json:"is_set_young"`
}

type BookInfoTemplate struct {
	BookId          string `json:"book_id"`
	BookName        string `json:"book_name"`
	Description     string `json:"description"`
	BookSrc         string `json:"book_src"`
	CategoryIndex   string `json:"category_index"`
	Tag             string `json:"tag"`
	TotalWordCount  string `json:"total_word_count"`
	UpStatus        string `json:"up_status"`
	UpdateStatus    string `json:"update_status"`
	IsPaid          string `json:"is_paid"`
	Discount        string `json:"discount"`
	DiscountEndTime string `json:"discount_end_time"`
	Cover           string `json:"cover"`
	AuthorName      string `json:"author_name"`
	Uptime          string `json:"uptime"`
	Newtime         string `json:"newtime"`
	ReviewAmount    string `json:"review_amount"`
	RewardAmount    string `json:"reward_amount"`
	ChapterAmount   string `json:"chapter_amount"`
	IsOriginal      string `json:"is_original"`
	TotalClick      string `json:"total_click"`
	MonthClick      string `json:"month_click"`
	WeekClick       string `json:"week_click"`
	MonthNoVipClick string `json:"month_no_vip_click"`
	WeekNoVipClick  string `json:"week_no_vip_click"`
	TotalRecommend  string `json:"total_recommend"`
	MonthRecommend  string `json:"month_recommend"`
	WeekRecommend   string `json:"week_recommend"`
	TotalFavor      string `json:"total_favor"`
	MonthFavor      string `json:"month_favor"`
	WeekFavor       string `json:"week_favor"`
	CurrentYp       string `json:"current_yp"`
	TotalYp         string `json:"total_yp"`
	CurrentBlade    string `json:"current_blade"`
	TotalBlade      string `json:"total_blade"`
	WeekFansValue   string `json:"week_fans_value"`
	MonthFansValue  string `json:"month_fans_value"`
	TotalFansValue  string `json:"total_fans_value"`
	LastChapterInfo struct {
		ChapterId         string `json:"chapter_id"`
		BookId            string `json:"book_id"`
		ChapterIndex      string `json:"chapter_index"`
		ChapterTitle      string `json:"chapter_title"`
		Uptime            string `json:"uptime"`
		Mtime             string `json:"mtime"`
		RecommendBookInfo string `json:"recommend_book_info"`
	} `json:"last_chapter_info"`
	TagList []struct {
		TagId   string `json:"tag_id"`
		TagType string `json:"tag_type"`
		TagName string `json:"tag_name"`
	} `json:"tag_list"`
	BookType        string `json:"book_type"`
	TransverseCover string `json:"transverse_cover"`
	GloryTag        struct {
		TagName    string `json:"tag_name"`
		CornerName string `json:"corner_name"`
		LabelIcon  string `json:"label_icon"`
		LinkUrl    string `json:"link_url"`
	} `json:"glory_tag"`
}
type CatalogueInfoTemplate struct {
	ChapterId      string `json:"chapter_id"`
	ChapterIndex   string `json:"chapter_index"`
	ChapterTitle   string `json:"chapter_title"`
	WordCount      string `json:"word_count"`
	TsukkomiAmount string `json:"tsukkomi_amount"`
	IsPaid         string `json:"is_paid"`
	Mtime          string `json:"mtime"`
	IsValid        string `json:"is_valid"`
	AuthAccess     string `json:"auth_access"`
}
type CatalogueInfoListTemplate struct {
	ChapterList     []CatalogueInfoTemplate `json:"chapter_list"`
	MaxUpdateTime   string                  `json:"max_update_time"`
	MaxChapterIndex string                  `json:"max_chapter_index"`
	DivisionId      string                  `json:"division_id"`
	DivisionIndex   string                  `json:"division_index"`
	DivisionName    string                  `json:"division_name"`
}
