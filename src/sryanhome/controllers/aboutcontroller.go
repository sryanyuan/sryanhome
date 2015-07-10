package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Get() {
	this.Layout = "layout.html"
	this.TplNames = "about/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "footer.html"
	this.LayoutSections["nav"] = "nav.html"
}
