package controllers

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

func (c *ArticleController) Index() {

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "article/article_headcssjs.html"
	c.LayoutSections["footerjs"] = "article/article_footerjs.html"
}
