package config

import (
	"net/http"

	admin "go_blog/admin/controllers"
	site "go_blog/site/controllers"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//admin
	//blog posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.New_Item)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Categories
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	//userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login/", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//Site
	r.GET("/", site.Homepage{}.Index)
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)

	//serve files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
