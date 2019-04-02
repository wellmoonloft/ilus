package utils

type Settings struct {
	General []general
	Seo     []seo
	Article []article
	Comment []comment
	Attach  []attach
	Backend []backend
	Email   []email
	Other   []other
}

//general 常规设置
type general struct {
	Language []language
	Title    string
	Url      string
	Logo     string
	Favicon  string
}

//Language 语言设置
type language struct {
	Id   int
	Name string
	Code string
}

//Seo SEO设置
type seo struct {
	KeyWord       []string
	Abs           string
	BaiDuToken    string
	GoogleCaptcha string
	BingCaptcha   string
	BaiDuCaptcha  string
}

//Article 文章设置
type article struct {
	IndexNum int
	RssNum   int
	AbsNum   int
}

//Comment 评论设置
type comment struct {
	Aevter          []aevter
	IsShowExamine   bool
	IsNewInform     bool
	IsExamineInform bool
	IsReplyInform   bool
	IsAPI           bool
	IndexNum        int
}

//Aevter 头像设置
type aevter struct {
	Id   int
	Name string
	Url  string
}

//Attach 附件设置
type attach struct {
	Position         int
	PositionSettings []positionSettings
}

//PositionSettings 存储位置设置
type positionSettings struct {
	Id        int
	Name      string
	Url       string
	AccessKey string
	SecretKey string
	Bucket    string
	Strategy  string
}

//Backend 后台设置
type backend struct {
	IsPjax    bool
	Layout    int
	IsSidebar bool
}

//Email 邮件设置
type email struct {
	IsEmail  bool
	Smtp     string
	UserName string
	PassWord string
	SendName string
}

//Other 其他设置
type other struct {
	IsAPI bool
	Token string
}
