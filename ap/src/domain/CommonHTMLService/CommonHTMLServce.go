package CommonHTMLService

import (
	"alma-server/ap/src/common/util/htmlutil"
	"html/template"
)

// よく使われるパーツの取得

// GetHead .
func GetHead() template.HTML {
	return template.HTML(htmlutil.CreateTemplateToString("/template/common/head.html", ""))
}

// GetHeader .
func GetHeader() template.HTML {
	return template.HTML(htmlutil.CreateTemplateToString("/template/common/header.html", ""))
}

// GetFooter .
func GetFooter() template.HTML {
	return template.HTML(htmlutil.CreateTemplateToString("/template/common/footer.html", ""))
}
