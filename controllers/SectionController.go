package controllers

import (
  "github.com/astaxie/beego"
  "pybbs-go/filters"
  "pybbs-go/models"
  "strconv"
)

type SectionController struct {
  beego.Controller
}

func (c *SectionController) List() {
  c.Data["PageTitle"] = "版块列表"
  c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
  c.Data["Sections"] = models.FindAllSection()
  c.Layout = "layout/layout.tpl"
  c.TplName = "section/list.tpl"
}

func (c *SectionController) Add() {
  beego.ReadFromRequest(&c.Controller)
  c.Data["PageTitle"] = "添加版块"
  c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
  c.Layout = "layout/layout.tpl"
  c.TplName = "section/add.tpl"
}

func (c *SectionController) Save() {
  flash := beego.NewFlash()
  name := c.GetString("name")
  if len(name) == 0 {
    flash.Error("版块名称不能为空")
    flash.Store(&c.Controller)
    c.Redirect("/section/add", 302)
  } else {
    section := models.Section{Name: name}
    models.SaveSection(&section)

    c.Redirect("/section/list", 302)
  }
}

func (c *SectionController) Edit() {
  beego.ReadFromRequest(&c.Controller)
  c.Data["PageTitle"] = "编辑版块"
  c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
  id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  if id > 0 {
    c.Data["Section"] = models.FindSectionById(id)
    c.Layout = "layout/layout.tpl"
    c.TplName = "section/edit.tpl"
  } else {
    c.Ctx.WriteString("版块不存在")
  }
}

func (c *SectionController) Update() {
  flash := beego.NewFlash()
  id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  name := c.GetString("name")
  if len(name) == 0 {
    flash.Error("版块名称不能为空")
    flash.Store(&c.Controller)
    c.Redirect("/section/add", 302)
  } else {
    section := models.Section{Id: id, Name: name}
    models.UpdateSection(&section)
    
    c.Redirect("/section/list", 302)
  }
}

func (c *SectionController) Delete() {
  id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  if id > 0 {
    section := models.Section{Id: id}
    models.DeleteSection(&section)
    c.Redirect("/section/list", 302)
  } else {
    c.Ctx.WriteString("版块不存在")
  }
}
