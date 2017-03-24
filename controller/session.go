package controller

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
)

var key = "bookshop_session"

func check(ctx *iris.Context) bool {

	// Check if user is authenticated
	if auth, _ := ctx.Session().GetBoolean("authenticated"); !auth {
		return false
	} else {
		return true
	}
}

func login(ctx *iris.Context) {
	session := ctx.Session()
	session.Set("authenticated", true)
}

func logout(ctx *iris.Context) {
	session := ctx.Session()
	session.Set("authenticated", false)
	ctx.Redirect("/", 302)
}

var sess = sessions.New(sessions.Config{Cookie: key})
