package controllers

import (
	"encoding/json"
	"go_practice/models"

	"github.com/astaxie/beego"
	"log"
)

// Operations about Users
type EtcdController struct {
	beego.Controller
}

// @Title Set
// @Description create Etcd
// @Param	body		body 	models.Etcd	true		"body for Etcd content"
// @Success 200  成功
// @Failure 403  失败
// @router / [post]
func (u *EtcdController) Set() {
	var etcd models.Etcd
	json.Unmarshal(u.Ctx.Input.RequestBody, &etcd)
	err := models.SetEtcd(etcd.Key, etcd.Value)
	if err != nil {
		u.Data["json"] = "失败"
	} else {
		u.Data["json"] = "成功"
	}
	u.ServeJSON()
}

// @Title Get
// @Description get etcd by k
// @Param	k		path 	string	true		"The key for staticblock"
// @Success 200 "data"
// @Failure 403 ""
// @router /:k [get]
func (u *EtcdController) Get() {
	k := u.GetString(":k")
	data := models.GetEtcd(k)
	log.Print(data)
	u.Data["json"] = data
	u.ServeJSON()
}
