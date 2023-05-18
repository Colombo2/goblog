package controllers

import (
	"fmt"
	"go_blog/admin/helpers"
	"go_blog/admin/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post_model{}.Get_All()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) New_Item(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.Get_All()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	r.ParseMultipartForm(10 << 20)
	file, header, erro := r.FormFile("blog-picture")
	if erro != nil {
		fmt.Println(erro)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, erroR := io.Copy(f, file)
	if erroR != nil {
		fmt.Println(erroR)
		return
	}

	models.Post_model{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Pic_url:     "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Kayıt Başarılı")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	post := models.Post_model{}.Get(params.ByName("id"))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post_model{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.Get_All()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if helpers.CheckUser(w, r) != true {
		helpers.SetAlert(w, r, "Lütfen Giriş Yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	post := models.Post_model{}.Get("id")
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")

	var pic_url string

	if is_selected == "1" {

		r.ParseMultipartForm(10 << 20)
		file, header, erro := r.FormFile("blog-picture")
		if erro != nil {
			fmt.Println(erro)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, erroR := io.Copy(f, file)
		if erroR != nil {
			fmt.Println(erroR)
			return
		}
		pic_url = "uploads/" + header.Filename
		os.Remove(post.Pic_url)
	} else {
		pic_url = post.Pic_url
	}

	post.Updates(models.Post_model{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Pic_url:     pic_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
