package routers

import (
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func AdminPassword(r render.Render, s sessions.Session) {
	r.HTML(200, "admin/password", map[string]interface{}{
		"IsPassword": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminUpdatePassword(req *http.Request, r render.Render, db *mgo.Database) {
	pass := req.FormValue("password")
	verify := req.FormValue("verify")

	if pass == verify {

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		info, err := db.C("id").Upsert(bson.M{"email": "admin@vim-tips.com"}, bson.M{"$set": bson.M{"password": hashedPass, "email": "admin@vim-tips.com"}})

		if err == nil {
			if info.Updated >= 0 {
				r.HTML(200, "admin/password", map[string]interface{}{
					"IsPost":     true,
					"IsPassword": true,
					"IsSuccess":  true}, render.HTMLOptions{Layout: "admin/layout"})
			}
		}
	} else {
		r.HTML(200, "admin/password", map[string]interface{}{
			"IsPost":     true,
			"IsPassword": true,
			"IsSuccess":  false}, render.HTMLOptions{Layout: "admin/layout"})
	}
}
