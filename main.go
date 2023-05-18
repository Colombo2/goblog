package main

import (
	"net/http"

	admin_models "go_blog/admin/models"
	"go_blog/config"
)

func main() {
	admin_models.Post_model{}.Migrate()
	admin_models.User_model{}.Migrate()
	admin_models.Category{}.Migrate()
	http.ListenAndServe("localhost:8080", config.Routes())
}
