package routers

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/mojo-zd/vim-tips-web/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func validateSession(r render.Render, s sessions.Session) {
	isLogin := s.Get("IsLogin")

	if isLogin == nil {
		fmt.Println("Not login...")
		r.Redirect("/admin/login")
	}
}

func ShowLoginPage(r render.Render, s sessions.Session) {

	isLogin := s.Get("IsLogin")

	if isLogin != nil {
		r.Redirect("/admin/index")
		return
	}

	r.HTML(200, "admin/login", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{})
}

func HandleLogout(r render.Render, s sessions.Session) {
	s.Delete("IsLogin")
	r.Redirect("/admin/login")
}

func HandleLogin(req *http.Request, r render.Render, s sessions.Session, db *mgo.Database) {
	username := req.FormValue("username")
	pass := req.FormValue("password")

	id := models.Identity{}

	err := db.C("id").Find(bson.M{"email": "admin@vim-tips.com"}).One(&id)

	if err != nil {
		fmt.Println(err.Error())
	}

	if username == "admin@vim-tips.com" && bcrypt.CompareHashAndPassword(id.Password, []byte(pass)) == nil {
		fmt.Println("Login success!")
		s.Set("IsLogin", true)

		r.Redirect("/admin/index")
	} else {
		r.Redirect("/admin/login")
	}
}

func HandleAdminIndex(r render.Render, db *mgo.Database) {
	tips_count, err := db.C("tips").Count()

	if err != nil {
		tips_count = 0
	}

	casts_count, err := db.C("casts").Count()

	if err != nil {
		casts_count = 0
	}

	r.HTML(200, "admin/index", map[string]interface{}{
		"IsIndex":    true,
		"TipsCount":  tips_count,
		"CastsCount": casts_count}, render.HTMLOptions{Layout: "admin/layout"})
}
