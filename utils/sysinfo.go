package utils

type Settings struct {
	General
	Seo
	Article
	Comment
	Attach
	Backend
	Email
	Other
}

//General 常规设置
type General struct {
	Language
	title   string
	url     string
	logo    string
	favicon string
}

//Language 语言设置
type Language struct {
	id   int
	name string
	code string
}

//Seo SEO设置
type Seo struct {
	keyWord       []string
	abs           string
	baiDuToken    string
	GoogleCaptcha string
	bingCaptcha   string
	baiDuCaptcha  string
}

//Article 文章设置
type Article struct {
	indexNum int
	rssNum   int
	absNum   int
}

//Comment 评论设置
type Comment struct {
	Aevter
	isShowExamine   bool
	isNewInform     bool
	isExamineInform bool
	isReplyInform   bool
	isAPI           bool
	indexNum        int
}

//Aevter 头像设置
type Aevter struct {
	id   int
	name string
	url  string
}

//Attach 附件设置
type Attach struct {
	position int
	PositionSettings
}

//PositionSettings 存储位置设置
type PositionSettings struct {
	id        int
	name      string
	url       string
	accessKey string
	secretKey string
	bucket    string
	strategy  string
}

//Backend 后台设置
type Backend struct {
	isPjax    bool
	layout    int
	isSidebar bool
}

//Email 邮件设置
type Email struct {
	isEmail  bool
	smtp     string
	userName string
	passWord string
	sendName string
}

//Other 其他设置
type Other struct {
	isAPI bool
	token string
}
