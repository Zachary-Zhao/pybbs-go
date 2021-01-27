package models

import (
  "github.com/astaxie/beego/orm"
)

type Section struct {
  Id   int    `orm:"pk;auto"`
  Name string `orm:"unique"`
}

func FindSectionById(id int) Section {
  o := orm.NewOrm()
  var sec Section
  o.QueryTable(sec).Filter("Id", id).One(&sec)
  return sec
}

func FindAllSection() []*Section {
  o := orm.NewOrm()
  var section Section
  var sections []*Section
  o.QueryTable(section).All(&sections)
  return sections
}

func SaveSection(sec *Section) int64 {
  o := orm.NewOrm()
  id, _ := o.Insert(sec)
  return id
}

func UpdateSection(sec *Section) {
  o := orm.NewOrm()
  o.Update(sec, "Name")
}

func DeleteSection(sec *Section) {
	o := orm.NewOrm()
	o.Delete(sec)
}