package catapi

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