package controller

import (
	"log"

	"strings"

	"github.com/atakanozceviz/bookshopui/model"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func Start() *iris.Framework {
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		view.HTML("./view/templates/", ".html"),
		sess,
	)
	app.StaticWeb("/public/", "./view/public/")
	app.Get("/", loginPage)
	app.Post("/girisYap", girisYap)
	app.Get("/cikisYap", logout)
	app.Get("/index", index)
	app.Get("/crud", crud)
	app.Get("/search/*param", search)
	return app
}

func loginPage(ctx *iris.Context) {
	if check(ctx) {
		ctx.Redirect("index", 307)
	} else {
		ctx.MustRender("login.html", model.PageData{
			Title: "Giriş Sayfası",
		})
	}
}

func girisYap(ctx *iris.Context) {
	var u model.User
	err := ctx.ReadForm(&u)
	if err != nil {
		log.Println(err)
	}
	if u.KSifre == "" || u.KAdi == "" {
		ctx.WriteString("empty")
	} else {
		id := u.CheckUser(u)
		if id != 0 {
			login(ctx)
			ctx.WriteString("success")
		} else {
			ctx.WriteString("fail")
		}
	}
}

func index(ctx *iris.Context) {
	if !check(ctx) {
		ctx.Redirect("/", 307)
	} else {
		ctx.MustRender("index.html", model.PageData{
			Title:  "Yönetici Paneli",
			Active: "index",
			User:   "Atakan Özceviz",
			Type:   "Admin",
		})
	}
}

func crud(ctx *iris.Context) {
	if !check(ctx) {
		ctx.Redirect("/", 307)
	} else {
		ctx.MustRender("crud.html", model.PageData{
			Title:  "Veri Tabanı",
			Active: "crud",
			User:   "Atakan Özceviz",
			Type:   "Admin",
		})
	}
}

func search(ctx *iris.Context) {
	if !check(ctx) {
		ctx.Redirect("/", 307)
	} else {
		var res model.Books
		k := ctx.URLParam("search")
		if k != "" {
			err := GetJSON("https://scrapeer.herokuapp.com/?keyword="+k, &res)
			if err != nil {
				log.Println(err)
			}
			r := strings.NewReplacer("136x136-0", "500x400-0")
			for i := range res {
				res[i].Img = r.Replace(res[i].Img)
			}
		}
		ctx.MustRender("search.html", model.PageData{
			Title:  "Kitap Arama",
			Active: "search",
			User:   "Atakan Özceviz",
			Type:   "Admin",
			Books:  res,
		})
	}
}
