package controller

import (
	"log"

	"strings"

	"github.com/atakanozceviz/bookShopUi/model"
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
	)
	app.StaticWeb("/public/", "./view/public/")
	app.Get("/", index)
	app.Get("/crud", crud)
	app.Get("/search/*param", search)
	return app
}

func index(ctx *iris.Context) {
	ctx.MustRender("index.html", model.PageData{
		Title:  "Yönetici Paneli",
		Active: "index",
		User:   "Atakan Özceviz",
		Type:   "Admin",
	})
}

func crud(ctx *iris.Context) {
	ctx.MustRender("crud.html", model.PageData{
		Title:  "Veri Tabanı",
		Active: "crud",
		User:   "Atakan Özceviz",
		Type:   "Admin",
	})
}

func search(ctx *iris.Context) {
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
