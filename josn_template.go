package catapi

type BookInfoQuery struct {
	UseDaguan        string `json:"use_daguan"`
	ModuleId         string `json:"module_id"`
	TabType          string `json:"tab_type"`
	Recommend        string `json:"recommend"`
	CarouselPosition string `json:"carousel_position"`
	BookId           string `json:"book_id"`
}
type ShelfListQuery struct {
	ShelfId     string `json:"shelf_id"`
	LastModTime string `json:"last_mod_time"`
	Direction   string `json:"direction"`
	Order       string `json:"order"`
	Count       string `json:"count"`
	Page        string `json:"page"`
}
type SearchTagFilterQuery struct {
	Filter string `json:"filter"`
	Tag    string `json:"tag"`
}
type SearchKeywordQuery struct {
	Count         string `json:"count"`
	Page          string `json:"page"`
	CategoryIndex string `json:"category_index"`
	Key           string `json:"key"`
}

type SearchTagsQuery struct {
	FilterWord    interface{}            `json:"filter_word"`
	Count         string                 `json:"count"`
	UseDaguan     string                 `json:"use_daguan"`
	Page          interface{}            `json:"page"`
	IsPaid        string                 `json:"is_paid"`
	CategoryIndex string                 `json:"category_index"`
	Key           string                 `json:"key"`
	FilterUptime  string                 `json:"filter_uptime"`
	UpStatus      string                 `json:"up_status"`
	Order         string                 `json:"order"`
	Tags          []SearchTagFilterQuery `json:"tags"`
}
type ContentInfoTemplate struct {
	ChapterId         string `json:"chapter_id" bson:"chapter_id"`
	BookId            string `json:"book_id" bson:"book_id"`
	DivisionId        string `json:"division_id" bson:"division_id"`
	UnitHlb           string `json:"unit_hlb" bson:"unit_hlb"`
	ChapterIndex      string `json:"chapter_index" bson:"chapter_index"`
	ChapterTitle      string `json:"chapter_title" bson:"chapter_title"`
	AuthorSay         string `json:"author_say" bson:"author_say"`
	WordCount         string `json:"word_count" bson:"word_count"`
	Discount          string `json:"discount" bson:"discount"`
	IsPaid            string `json:"is_paid" bson:"is_paid"`
	AuthAccess        string `json:"auth_access" bson:"auth_access"`
	BuyAmount         string `json:"buy_amount" bson:"buy_amount"`
	TsukkomiAmount    string `json:"tsukkomi_amount" bson:"tsukkomi_amount"`
	TotalHlb          string `json:"total_hlb" bson:"total_hlb"`
	Uptime            string `json:"uptime" bson:"uptime"`
	Mtime             string `json:"mtime" bson:"mtime"`
	Ctime             string `json:"ctime" bson:"ctime"`
	RecommendBookInfo string `json:"recommend_book_info" bson:"recommend_book_info"`
	BaseStatus        string `json:"base_status" bson:"base_status"`
	TxtContent        string `json:"txt_content" bson:"txt_content"`
}
type BookInfoTemplate struct {
	AuthorName      string `json:"author_name" bson:"author_name"`
	BookId          string `json:"book_id" bson:"book_id"`
	BookName        string `json:"book_name" bson:"book_name"`
	BookSrc         string `json:"book_src" bson:"book_src"`
	BookType        string `json:"book_type" bson:"book_type"`
	CategoryIndex   string `json:"category_index" bson:"category_index"`
	ChapterAmount   string `json:"chapter_amount" bson:"chapter_amount"`
	Cover           string `json:"cover" bson:"cover"`
	CurrentBlade    string `json:"current_blade" bson:"current_blade"`
	CurrentYp       string `json:"current_yp" bson:"current_yp"`
	Description     string `json:"description" bson:"description"`
	Discount        string `json:"discount" bson:"discount"`
	DiscountEndTime string `json:"discount_end_time" bson:"discount_end_time"`
	GloryTag        struct {
		CornerName string `json:"corner_name" bson:"corner_name"`
		LabelIcon  string `json:"label_icon" bson:"label_icon"`
		LinkUrl    string `json:"link_url" bson:"link_url"`
		TagName    string `json:"tag_name" bson:"tag_name"`
	} `json:"glory_tag" bson:"glory_tag"`
	IsOriginal      string `json:"is_original" bson:"is_original"`
	IsPaid          string `json:"is_paid" bson:"is_paid"`
	LastChapterInfo struct {
		BookId            string `json:"book_id" bson:"book_id"`
		ChapterId         string `json:"chapter_id" bson:"chapter_id"`
		ChapterIndex      string `json:"chapter_index" bson:"chapter_index"`
		ChapterTitle      string `json:"chapter_title" bson:"chapter_title"`
		Mtime             string `json:"mtime" bson:"mtime"`
		RecommendBookInfo string `json:"recommend_book_info" bson:"recommend_book_info"`
		Uptime            string `json:"uptime" bson:"uptime"`
	} `json:"last_chapter_info" bson:"last_chapter_info"`
	MonthClick      string `json:"month_click" bson:"month_click"`
	MonthFansValue  string `json:"month_fans_value" bson:"month_fans_value"`
	MonthFavor      string `json:"month_favor" bson:"month_favor"`
	MonthNoVipClick string `json:"month_no_vip_click" bson:"month_no_vip_click"`
	MonthRecommend  string `json:"month_recommend" bson:"month_recommend"`
	Newtime         string `json:"newtime" bson:"newtime"`
	ReviewAmount    string `json:"review_amount" bson:"review_amount"`
	RewardAmount    string `json:"reward_amount" bson:"reward_amount"`
	Tag             string `json:"tag" bson:"tag"`
	TagList         []struct {
		TagId   string `json:"tag_id" bson:"tag_id"`
		TagName string `json:"tag_name" bson:"tag_name"`
		TagType string `json:"tag_type" bson:"tag_type"`
	} `json:"tag_list" bson:"tag_list"`
	TotalBlade      string `json:"total_blade" bson:"total_blade"`
	TotalClick      string `json:"total_click" bson:"total_click"`
	TotalFansValue  string `json:"total_fans_value" bson:"total_fans_value"`
	TotalFavor      string `json:"total_favor" bson:"total_favor"`
	TotalRecommend  string `json:"total_recommend" bson:"total_recommend"`
	TotalWordCount  string `json:"total_word_count" bson:"total_word_count"`
	TotalYp         string `json:"total_yp" bson:"total_yp"`
	TransverseCover string `json:"transverse_cover" bson:"transverse_cover"`
	UpStatus        string `json:"up_status" bson:"up_status"`
	UpdateStatus    string `json:"update_status" bson:"update_status"`
	Uptime          string `json:"uptime" bson:"uptime"`
	WeekClick       string `json:"week_click" bson:"week_click"`
	WeekFansValue   string `json:"week_fans_value" bson:"week_fans_value"`
	WeekFavor       string `json:"week_favor" bson:"week_favor"`
	WeekNoVipClick  string `json:"week_no_vip_click" bson:"week_no_vip_click"`
	WeekRecommend   string `json:"week_recommend" bson:"week_recommend"`
}
type ChapterListTemplate struct {
	Code string      `json:"code" bson:"code"`
	Tip  interface{} `json:"tip" bson:"tip"`
	Data struct {
		ChapterList []struct {
			ChapterList []struct {
				ChapterId      string `json:"chapter_id" bson:"chapter_id"`
				ChapterIndex   string `json:"chapter_index" bson:"chapter_index"`
				ChapterTitle   string `json:"chapter_title" bson:"chapter_title"`
				WordCount      string `json:"word_count" bson:"word_count"`
				TsukkomiAmount string `json:"tsukkomi_amount" bson:"tsukkomi_amount"`
				IsPaid         string `json:"is_paid" bson:"is_paid"`
				Mtime          string `json:"mtime" bson:"mtime"`
				IsValid        string `json:"is_valid" bson:"is_valid"`
				AuthAccess     string `json:"auth_access" bson:"auth_access"`
			} `json:"chapter_list" bson:"chapter_list"`
			MaxUpdateTime   string `json:"max_update_time" bson:"max_update_time"`
			MaxChapterIndex string `json:"max_chapter_index" bson:"max_chapter_index"`
			DivisionId      string `json:"division_id" bson:"division_id"`
			DivisionIndex   string `json:"division_index" bson:"division_index"`
			DivisionName    string `json:"division_name" bson:"division_name"`
		} `json:"chapter_list" bson:"chapter_list"`
	} `json:"data" bson:"data"`
	ScrollChests []interface{} `json:"scroll_chests" bson:"scroll_chests"`
}
