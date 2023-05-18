package controllers

import (
	"fmt"
	"go_blog/admin/helpers"
	"go_blog/admin/models"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("usern")
	password := r.FormValue("passw")
	user := models.User_model{}.Get("username = ? AND 1", username)
	or_not := helpers.Compare(user.Password, []byte(password))
	if user.Username == username && or_not == true {
		helpers.SetUser(w, r, username, user.Password)
		helpers.SetAlert(w, r, "Hoşgeldiniz")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		helpers.SetAlert(w, r, "Hatalı Bilgi Girişi")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Hoşçakalın")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
