package controllers

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type MainController struct {
	beego.Controller
}

func NewClient() (client.Client, error) {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return c, err
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetKey() {
	key := c.GetString("key")
	cl, _ := NewClient()
	kapi := client.NewKeysAPI(cl)

	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get is done. Metadata is %q\n", resp)
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}

	c.Data["json"] = resp.Node.Value
	c.ServeJSON()
}

func (c *MainController) SetKey() {
	key := c.GetString("key")
	value := c.GetString("value")

	cl, _ := NewClient()
	kapi := client.NewKeysAPI(cl)

	resp, err := kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
	}

	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (c *MainController) DelKey() {
	key := c.GetString("key")
	cl, _ := NewClient()
	kapi := client.NewKeysAPI(cl)

	resp, err := kapi.Delete(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp.Node.Key)
	}

	c.Data["json"] = "ok"
	c.ServeJSON()
}
