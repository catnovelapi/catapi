package catapi

type BookInfoQuery struct {
	UseDaguan        int    `json:"use_daguan"`
	ModuleId         int    `json:"module_id"`
	TabType          int    `json:"tab_type"`
	Recommend        string `json:"recommend"`
	CarouselPosition string `json:"carousel_position"`
	BookId           string `json:"book_id"`
}
type SearchTagFilterQuery struct {
	Filter int    `json:"filter"`
	Tag    string `json:"tag"`
}
type SearchKeywordQuery struct {
	Count         int    `json:"count"`
	Page          string `json:"page"`
	CategoryIndex int    `json:"category_index"`
	Key           string `json:"key"`
}

type SearchTagsQuery struct {
	FilterWord    interface{}            `json:"filter_word"`
	Count         int                    `json:"count"`
	UseDaguan     int                    `json:"use_daguan"`
	Page          interface{}            `json:"page"`
	IsPaid        string                 `json:"is_paid"`
	CategoryIndex int                    `json:"category_index"`
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
